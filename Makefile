# Ecosystem images Makefile
# This makefile aims to help the developer and release captain

# git ls-remote ${URL} refs/heads/${BRANCH} | awk '{print $1 }'
IMAGES=$(shell yq e 'keys[]' sources.yaml)
IMAGES_UP=$(addprefix sources/, $(addsuffix /upgrade,$(IMAGES)))

.PHONY: actions/matrix
actions/matrix:
	@yq e 'keys | map_values("name": .) | flatten' sources.yaml  | yq e '{"include": .}' | yq -o json | jq --indent 0 -r

.PHONY: sources/upgrade
sources/upgrade: $(IMAGES_UP)

.PHONY: sources
sources: $(addprefix sources/,$(IMAGES))

sources/%/upgrade:
	./hack/upgrade-source.sh $*

sources/%:
	./hack/checkout-source.sh $*

