#!/usr/bin/env bash
set -e
test -e opt/fpm/.env && . opt/fpm/.env # Source local test env

pwd
ls -la

pkill ssh-agent||true
eval $(ssh-agent)
ssh-add - <<<"$REPO_SSH_PRIV_KEY"
test -e ~.ssh || mkdir -p ~/.ssh
grep -q "$REPO_SSH_HOST" ~/.ssh/known_hosts || echo "$REPO_SSH_KNOWN_HOSTS">>~/.ssh/known_hosts
echo "👷 Copying to repo server ... "
scp -P ${REPO_SSH_PORT} *.deb "${REPO_SSH_USER}"@"${REPO_SSH_HOST}":~/incoming/
scp -P ${REPO_SSH_PORT} *.rpm "${REPO_SSH_USER}"@"${REPO_SSH_HOST}":~/incoming/
echo "✅ All files copied"
echo "👷 Triggering package publishing ... "
ssh -p "${REPO_SSH_PORT}" "${REPO_SSH_USER}"@"${REPO_SSH_HOST}" bash <<"EOF"
cd ~/incomning
ls -la
DISTS="jammy focal bionic bullseye buster"
for DIST in $DISTS;do
   aptly repo add ${DIST}-unstable *.deb
done

## Update all repos, otherwise newly added packages are not published
for DIST in $DISTS;do
   aptly publish update --gpg-key="${REPO_GPG_KEY_ID}" $DIST
done
echo "==========================================================="
echo " Debian Repos Updated"
echo "==========================================================="

rpm --addsign ~/incoming/*.rpm
mv *.rpm ~/public/rpm/unstable
~/bin/rpm-updaterepo.sh
echo "==========================================================="
echo " RPM Repos Updated"
echo "==========================================================="
EOF
echo "✅ All packages published"