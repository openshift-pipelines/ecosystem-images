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
    pushd $TMPDIR/$GITOPSNAME
    git reset --hard HEAD
    git checkout main
    git fetch -p --all
    git rebase origin/main
    popd
fi

cp -r .konflux/* "$TMPDIR/$GITOPSNAME/$ECOSYSTEMPATH/$ECOSYSTEMIMAGES"
rm "$TMPDIR/$GITOPSNAME/$ECOSYSTEMPATH/$ECOSYSTEMIMAGES/README.md"
pushd "$TMPDIR/$GITOPSNAME/$ECOSYSTEMPATH"
kustomize edit remove resource "$ECOSYSTEMIMAGES/*"
kustomize edit remove resource "$ECOSYSTEMIMAGES/*/*"
find "$ECOSYSTEMIMAGES" -type f -print -exec kustomize edit add resource {} \;

cat <<EOF > $TMPDIR/yamlfmt.conf
formatter:
  type: basic
  include_document_start: true
EOF
yamlfmt -conf $TMPDIR/yamlfmt.conf kustomization.yaml

pushd ../../../../
./build-manifests.sh
popd

git checkout -B sync-tekton-ecosystem-images
git add $ECOSYSTEMIMAGES kustomization.yaml
git commit -s -m "Synchronize tekton-ecosystem ecosystem-images definitions"
git push -u --force origin sync-tekton-ecosystem-images

popd
