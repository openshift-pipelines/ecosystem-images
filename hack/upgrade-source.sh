#!/usr/bin/env bash
set -euo pipefail

IMAGE=${1:=}
GITHUB=$(yq e ".${IMAGE}.github" sources.yaml)
BRANCH=$(yq e ".${IMAGE}.branch" sources.yaml)
COMMIT=$(yq e ".${IMAGE}.commit" sources.yaml)

NEW_COMMIT=$(git ls-remote https://github.com/${GITHUB} refs/heads/${BRANCH} | awk '{print $1 }')

echo "${GITHUB}#${BRANCH}: ${COMMIT}"
if [[ "${COMMIT}" = "${NEW_COMMIT}" ]]; then
    exit 0
fi

echo "Updating sources.yaml: ${NEW_COMMIT}"
yq e -i ".${IMAGE}.commit = \"${NEW_COMMIT}\"" sources.yaml

make "sources/${IMAGE}"
