#!/usr/bin/env bash
set -e
set -o nounset
set -o pipefail

# e.g.,
# ./scripts/build.sh
# ./scripts/tests.e2e.sh ./build/avalanchego
# PROMETHEUS_PATH=/tmp/prometheus ./scripts/tests.e2e.sh ./build/avalanchego
if ! [[ "$0" =~ scripts/tests.e2e.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi



AVALANCHEGO_PATH="${1-}"
if [[ -z "${AVALANCHEGO_PATH}" ]]; then
  echo "Missing AVALANCHEGO_PATH argument!"
  echo "Usage: ${0} [AVALANCHEGO_PATH]" >> /dev/stderr
  exit 255
fi

PROMETHEUS_PATH="${PROMETHEUS_PATH:-""}" 

# Set the CGO flags to use the portable version of BLST
#
# We use "export" here instead of just setting a bash variable because we need
# to pass this flag to all child processes spawned by the shell.
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"
# While CGO_ENABLED doesn't need to be explicitly set, it produces a much more
# clear error due to the default value change in go1.20.
export CGO_ENABLED=1

GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)

#################################
# download avalanche-network-runner
# https://github.com/ava-labs/avalanche-network-runner
# TODO: migrate to upstream avalanche-network-runner
NETWORK_RUNNER_VERSION=1.6.0
DOWNLOAD_PATH=/tmp/avalanche-network-runner.tar.gz
DOWNLOAD_URL="https://github.com/ava-labs/avalanche-network-runner/releases/download/v${NETWORK_RUNNER_VERSION}/avalanche-network-runner_${NETWORK_RUNNER_VERSION}_${GOOS}_${GOARCH}.tar.gz"

rm -f ${DOWNLOAD_PATH}
rm -f /tmp/avalanche-network-runner

echo "downloading avalanche-network-runner ${NETWORK_RUNNER_VERSION} at ${DOWNLOAD_URL} to ${DOWNLOAD_PATH}"
curl --fail -L ${DOWNLOAD_URL} -o ${DOWNLOAD_PATH}

echo "extracting downloaded avalanche-network-runner"
tar xzvf ${DOWNLOAD_PATH} -C /tmp
/tmp/avalanche-network-runner -h

# installs prometheus to PROMETHEUS_PATH if it does not already exist.
function download_prometheus {
  local version=2.44.0
  local download_path=/tmp/prometheus.tar.gz
  local download_url="https://github.com/prometheus/prometheus/releases/download/v${version}/prometheus-${version}.${GOOS}-${GOARCH}.tar.gz"

  if [ ! -f "$PROMETHEUS_PATH" ]; then
    rm -f ${download_path}
    echo "downloading prometheus ${version} at ${download_url} to ${download_path}"
    curl --fail -L ${download_url} -o ${download_path}

    echo "extracting downloaded prometheus"
    tar xvf "${download_path}" -C /tmp prometheus-"${version}.${GOOS}-${GOARCH}"/prometheus --strip-components 1
  fi

  /tmp/prometheus --version
}

if [ -z "${PROMETHEUS_PATH}" ]; then
  download_prometheus
fi

GOPATH="$(go env GOPATH)"
PATH="${GOPATH}/bin:${PATH}"

#################################
echo "building e2e.test"
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo@v2.1.4
ACK_GINKGO_RC=true ginkgo build ./tests/e2e
./tests/e2e/e2e.test --help

#################################
# run "avalanche-network-runner" server
echo "launch avalanche-network-runner in the background"
/tmp/avalanche-network-runner \
server \
--log-level debug \
--port=":12342" \
--disable-grpc-gateway &
PID=${!}

#################################
echo "running e2e tests against the local cluster with ${AVALANCHEGO_PATH}"
./tests/e2e/e2e.test \
--ginkgo.v \
--log-level debug \
--network-runner-grpc-endpoint="0.0.0.0:12342" \
--network-runner-avalanchego-path="${AVALANCHEGO_PATH}" \
--network-runner-avalanchego-log-level="WARN" \
--network-runner-prometheus-path="${PROMETHEUS_PATH}" \
--test-keys-file=tests/test.insecure.secp256k1.keys \
&& EXIT_CODE=$? || EXIT_CODE=$?

kill "${PID}"

if [[ ${EXIT_CODE} -gt 0 ]]; then
  echo "FAILURE with exit code ${EXIT_CODE}"
  exit ${EXIT_CODE}
else
  echo "ALL SUCCESS!"
fi
