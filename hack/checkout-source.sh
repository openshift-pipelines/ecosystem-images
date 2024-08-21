#!/usr/bin/env bash
set -euo pipefail

IMAGE=${1:=}
GITHUB=$(yq e ".${IMAGE}.github" sources.yaml)
COMMIT=$(yq e ".${IMAGE}.commit" sources.yaml)
GITPATH=$(yq e ".${IMAGE}.path" sources.yaml)

TMPDIR=${TMPDIR:=}
if [ -z "$TMPDIR" ]; then
    TMPDIR="$(mktemp -d)"
    trap 'rm -rf -- "$TMPDIR"' EXIT
fi

# Checkout code
pushd ${TMPDIR}
git init
git fetch --depth 1 https://github.com/${GITHUB} ${COMMIT}
git checkout FETCH_HEAD
rm -fR .git
popd

# Clean-up image content
rm -fR ${IMAGE}

# Copy image content
cp -fR ${TMPDIR}/${GITPATH} ${IMAGE}

# Generate image content
# FIXME: generate more
cp -fR openshift/${IMAGE}/* ${IMAGE}/
