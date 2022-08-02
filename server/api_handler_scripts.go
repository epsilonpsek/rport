package chserver

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cloudradar-monitoring/rport/server/api"
	errors2 "github.com/cloudradar-monitoring/rport/server/api/errors"
	"github.com/cloudradar-monitoring/rport/server/api/jobs"
	"github.com/cloudradar-monitoring/rport/server/auditlog"
	"github.com/cloudradar-monitoring/rport/server/clients"
	"github.com/cloudradar-monitoring/rport/share/ws"
)

// handleExecuteScript handles GET /clients/{client_id}/scripts
func (al *APIListener) handleExecuteScript(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	cid := vars[routeParamClientID]
	if cid == "" {
		al.jsonErrorResponseWithTitle(w, http.StatusBadRequest, fmt.Sprintf("Missing %q route param.", routeParamClientID))
		return
	}

	execCmdInput := &api.ExecuteInput{}
	err := parseRequestBody(req.Body, &execCmdInput)
	if err != nil {
		al.jsonError(w, err)
		return
	}
	if execCmdInput.Script == "" {
		al.jsonErrorResponseWithTitle(w, http.StatusBadRequest, "Missing script body")
		return
	}

	decodedScriptBytes, err := base64.StdEncoding.DecodeString(execCmdInput.Script)
	if err != nil {
		al.jsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	execCmdInput.Command = string(decodedScriptBytes)

	execCmdInput.ClientID = cid
	execCmdInput.IsScript = true

	resp := al.handleExecuteCommand(req.Context(), w, execCmdInput)

	if resp != nil {
		al.auditLog.Entry(auditlog.ApplicationClientScript, auditlog.ActionExecuteStart).
			WithHTTPRequest(req).
			WithClientID(cid).
			WithRequest(execCmdInput).
			WithResponse(resp).
			WithID(resp.JID).
			Save()
	}
}

// handlePostMultiClientScript handles POST /scripts
func (al *APIListener) handlePostMultiClientScript(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	inboundMsg := new(jobs.MultiJobRequest)
	err := parseRequestBody(req.Body, inboundMsg)
	if err != nil {
		al.jsonError(w, err)
		return
	}

	errTitle, err := checkTargetingParams(inboundMsg)
	if err != nil {
		al.jsonErrorResponseWithError(w, http.StatusBadRequest, errTitle, err)
		return
	}

	err = al.enrichScriptInput(ctx, inboundMsg)
	if err != nil {
		al.jsonError(w, err)
		return
	}

	orderedClients, clientsInGroupsCount, err := al.makeOrderedClientsForScript(ctx, inboundMsg)
	if err != nil {
		al.jsonError(w, err)
		return
	}

	if len(orderedClients) == 0 {
		al.jsonErrorResponseWithTitle(w, http.StatusBadRequest, "no clients to execute the script for")
		return
	}

	inboundMsg.OrderedClients = orderedClients

	if !hasClientTags(inboundMsg) {
		errTitle := validateNonClientsTagTargeting(inboundMsg, clientsInGroupsCount, orderedClients, 2)
		if errTitle != "" {
			al.jsonErrorResponseWithTitle(w, http.StatusBadRequest, errTitle)
			return
		}
	} else {
		errTitle := validateClientTagsTargeting(orderedClients)
		if errTitle != "" {
			al.jsonErrorResponseWithTitle(w, http.StatusBadRequest, errTitle)
			return
		}
	}

	curUser, err := al.getUserModelForAuth(req.Context())
	if err != nil {
		al.jsonError(w, err)
		return
	}

	err = al.clientService.CheckClientsAccess(inboundMsg.OrderedClients, curUser)
	if err != nil {
		al.jsonError(w, err)
		return
	}

	inboundMsg.Username = curUser.Username

	multiJob, err := al.StartMultiClientJob(ctx, inboundMsg)
	if err != nil {
		al.jsonError(w, err)
		return
	}

	resp := newJobResponse{
		JID: multiJob.JID,
	}

	al.auditLog.Entry(auditlog.ApplicationClientScript, auditlog.ActionExecuteStart).
		WithHTTPRequest(req).
		WithRequest(inboundMsg).
		WithResponse(resp).
		WithID(multiJob.JID).
		SaveForMultipleClients(inboundMsg.OrderedClients)

	al.writeJSONResponse(w, http.StatusOK, api.NewSuccessPayload(resp))

	al.Debugf("Multi-client Job[id=%q] created to execute remote command on clients %s, groups %s, tags %s: %q.", multiJob.JID, inboundMsg.ClientIDs, inboundMsg.GroupIDs, inboundMsg.GetTags(), inboundMsg.Command)
}

// handleScriptsWS handles GET /ws/scripts
func (al *APIListener) handleScriptsWS(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	uiConn, err := apiUpgrader.Upgrade(w, req, nil)
	if err != nil {
		al.Errorf("Failed to establish WS connection: %v", err)
		return
	}

	uiConnTS := ws.NewConcurrentWebSocket(uiConn, al.Logger)

	inboundMsg := &jobs.MultiJobRequest{}
	err = uiConnTS.ReadJSON(inboundMsg)
	if err == io.EOF { // is handled separately to return an informative error message
		uiConnTS.WriteError("Inbound message should contain non empty json object with command data.", nil)
		return
	}
	if err != nil {
		uiConnTS.WriteError("Invalid JSON data.", err)
		return
	}

	errTitle, err := checkTargetingParams(inboundMsg)
	if err != nil {
		uiConnTS.WriteError(errTitle, err)
		return
	}

	err = al.enrichScriptInput(ctx, inboundMsg)
	if err != nil {
		uiConnTS.WriteError("Failed to create script on multiple clients", err)
		return
	}

	orderedClients, clientsInGroupsCount, err := al.makeOrderedClientsForScript(ctx, inboundMsg)
	if err != nil {
		uiConnTS.WriteError("Failed to get client list", err)
		return
	}

	if len(orderedClients) == 0 {
		uiConnTS.WriteError("no clients to execute the script for", nil)
		return
	}

	inboundMsg.OrderedClients = orderedClients

	auditLogEntry := al.auditLog.Entry(auditlog.ApplicationClientScript, auditlog.ActionExecuteStart).WithHTTPRequest(req)

	al.handleCommandsExecutionWS(ctx, uiConnTS, inboundMsg, clientsInGroupsCount, auditLogEntry)
}

func (al *APIListener) enrichScriptInput(
	ctx context.Context,
	inboundMsg *jobs.MultiJobRequest,
) (err error) {
	if inboundMsg.Script == "" {
		return errors2.APIError{
			Message:    "Missing script body",
			HTTPStatus: http.StatusBadRequest,
		}
	}

	if inboundMsg.TimeoutSec <= 0 {
		inboundMsg.TimeoutSec = al.config.Server.RunRemoteCmdTimeoutSec
	}

	decodedScriptBytes, err := base64.StdEncoding.DecodeString(inboundMsg.Script)
	if err != nil {
		return errors2.APIError{
			Err:        err,
			HTTPStatus: http.StatusBadRequest,
			Message:    "failed to decode script payload from base64",
		}
	}

	inboundMsg.Command = string(decodedScriptBytes)
	inboundMsg.IsScript = true

	return nil
}

func (al *APIListener) makeOrderedClientsForScript(ctx context.Context, inboundMsg *jobs.MultiJobRequest) (orderedClients []*clients.Client, clientsInGroupsCount int, err error) {
	if !hasClientTags(inboundMsg) {
		orderedClients, clientsInGroupsCount, err = al.getOrderedClients(ctx, inboundMsg.ClientIDs, inboundMsg.GroupIDs, false /* allowDisconnected */)
		if err != nil {
			return nil, 0, err
		}
	} else {
		orderedClients, err = al.getOrderedClientsByTag(ctx, inboundMsg.ClientIDs, inboundMsg.GroupIDs, inboundMsg.ClientTags, false /* allowDisconnected */)
		if err != nil {
			return nil, 0, err
		}
	}

	return orderedClients, clientsInGroupsCount, nil
}
