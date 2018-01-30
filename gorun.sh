#!/bin/bash

export PROJ="github.com/gileshuang/gollector/cmd/gollectord"

export GOPATH=$(pwd)
export GOBIN=$(pwd)/bin

if [[ "x${PROJ}" == "x" ]]; then
    if [[ "x$1" != "x" ]]; then
        export PROJ=$1
    fi
fi

for DIR in $(echo ${GOPATH} | sed 's/\:/\ /'); do
    if [[ -d "${DIR}/src/${PROJ}" ]]; then
        export PROJ_DIR=${DIR}/src/$(echo ${PROJ} | awk -F '/' '{ print $1"/"$2"/"$3 }')
        break
    fi
done
if [[ "x${PROJ_DIR}" == "x" || ! -d "${PROJ_DIR}" ]]; then
    echo "Can not find project dir for ${PROJ}"
    exit 1
fi

GORUN_DIR=$(mktemp -d -t gorun.XXXXXXXX)

rsync -r "${PROJ_DIR}/" "${GORUN_DIR}/"

rm -f ${GOBIN}/`basename ${PROJ}`
go install ${PROJ}

cd "${GORUN_DIR}"
${GOBIN}/`basename ${PROJ}`

cd -
echo "clean GORUN_DIR"
rm -rf "${GORUN_DIR}"
