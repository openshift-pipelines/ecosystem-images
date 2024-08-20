#!/usr/bin/env bash
# This script synchronize this repository `.konflux` folder with
# the Red Hat instance GitOps repository.

set -eu

ECOSYSTEMIMAGES="ecosystem-images"
ECOSYSTEMPATH="tenants-config/cluster/stone-prd-rh01/tenants/tekton-ecosystem-tenant"
GITOPSNAME="konflux-release-data"
GITOPSREPOSITORY=$1
TMPDIR=${TMPDIR:=}
shift

if [ -z "$TMPDIR" ]; then
    TMPDIR="$(mktemp -d)"
    trap 'rm -rf -- "$TMPDIR"' EXIT
fi

if [ ! -d "$TMPDIR/$GITOPSNAME" ]; then
    echo "Cloning $GITOPSREPOSITORY in $TMPDIR"
    git clone "$GITOPSREPOSITORY" "$TMPDIR/$GITOPSNAME"
else
    echo "Making sure $GITOPSREPOSITORY in $TMPDIR is up-to-date"
    git reset --hard HEAD
    git checkout main
    git fetch -p --all
    git rebase origin/main
fi

cp -r .konflux/* "$TMPDIR/$GITOPSNAME/$ECOSYSTEMPATH/$ECOSYSTEMIMAGES"
cd "$TMPDIR/$GITOPSNAME/$ECOSYSTEMPATH"
kustomize edit remove resource "$ECOSYSTEMIMAGES/*"
kustomize edit remove resource "$ECOSYSTEMIMAGES/*/*"
find "$ECOSYSTEMIMAGES" -type f -print -exec kustomize edit add resource {} \;

git checkout -B sync-tekton-ecosystem-images
git add $ECOSYSTEMIMAGES kustomization.yaml
git commit -s -m "Synchronize tekton-ecosystem ecosystem-images definitions"
git push -u --force origin sync-tekton-ecosystem-images
