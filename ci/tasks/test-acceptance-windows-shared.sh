#!/bin/bash -eux

set -o pipefail

ROOT_DIR=$PWD
BBL_STATE_DIR=$ROOT_DIR/bbl-state

export PATH=$BBL_STATE_DIR/bin:$PATH
source $BBL_STATE_DIR/bosh.sh

bosh -n upload-stemcell $ROOT_DIR/bosh-candidate-stemcell-windows/*.tgz
bosh -n upload-stemcell $ROOT_DIR/gcp-linux-stemcell/*.tgz

pushd $ROOT_DIR/dns-release
   bosh create-release --force && bosh upload-release --rebase
popd

bosh -n -d bosh-dns-shared-acceptance deploy $ROOT_DIR/dns-release/src/test_yml_assets/shared-acceptance-manifest.yml \
    --var-file bosh_ca_cert=<(echo "$BOSH_CA_CERT") \
    -v bosh_client_secret="$BOSH_CLIENT_SECRET" \
    -v bosh_client="$BOSH_CLIENT" \
    -v bosh_environment="$BOSH_ENVIRONMENT" \
    -v bosh_deployment=bosh-dns

pushd $ROOT_DIR/dns-release/src/acceptance_tests/dns-acceptance-release
   bosh create-release --force && bosh upload-release --rebase
popd

bosh -d bosh-dns-shared-acceptance run-errand acceptance-tests
