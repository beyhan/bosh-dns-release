#!/bin/bash -eux

set -eu -o pipefail

ROOT_DIR=$PWD
BBL_STATE_DIR=$ROOT_DIR/envs/$ENV_NAME

pushd $BBL_STATE_DIR
  source .envrc
popd

bosh -n -d bosh-dns-shared-acceptance \
  deploy $ROOT_DIR/bosh-dns-release/src/bosh-dns/test_yml_assets/manifests/shared-acceptance-manifest.yml \
  --var-file bosh_ca_cert=<(echo "$BOSH_CA_CERT") \
  -v bosh_client_secret="$BOSH_CLIENT_SECRET" \
  -v bosh_client="$BOSH_CLIENT" \
  -v bosh_environment="$BOSH_ENVIRONMENT" \
  -v base_stemcell=$WINDOWS_OS_VERSION \
  -v bosh_deployment=bosh-dns \
  -v jumpbox_private_key="$(cat $JUMPBOX_PRIVATE_KEY)" \
  -v jumpbox_address="$JUMPBOX_ADDRESS" \
  --vars-store dns-creds.yml

pushd $ROOT_DIR/bosh-dns-release/src/bosh-dns/acceptance_tests/dns-acceptance-release
   bosh create-release --force && bosh upload-release --rebase
popd

bosh -d bosh-dns-shared-acceptance run-errand acceptance-tests --keep-alive
