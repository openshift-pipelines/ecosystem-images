FROM brew.registry.redhat.io/rh-osbs/openshift-golang-builder:v1.22 AS builder
# FROM registry.access.redhat.com/ubi9/go-toolset:1.21.10-1.1719562237@sha256:ae17d73e70a966f39ef4dfca74241e3ca4374cd1198b02c30ea0748b8dcc83a6 AS builder

COPY git-init .
ENV GODEBUG="http2server=0"
RUN go build -ldflags="-X 'knative.dev/pkg/changeset.rev=${CHANGESET_REV:0:7}'" -mod=vendor -v -o /tmp/tektoncd-catalog-git-clone 

FROM registry.redhat.io/ubi8/ubi-minimal:latest

ENV BINARY=git-init \
    KO_APP=/ko-app \
    KO_DATA_PATH=/kodata

RUN microdnf install -y openssh-clients git git-lfs shadow-utils

COPY --from=builder /tmp/tektoncd-catalog-git-clone ${KO_APP}/${BINARY}

RUN chgrp -R 0 ${KO_APP} && \
    chmod -R g+rwX ${KO_APP}

# version="${CI_CONTAINER_VERSION}" \
LABEL \
    com.redhat.component="openshift-pipelines-git-init-rhel8-container" \
    name="openshift-pipelines/pipelines-git-init-rhel8" \
    summary="Red Hat OpenShift Pipelines Git-init" \
    maintainer="pipelines-extcomm@redhat.com" \
    description="Red Hat OpenShift Pipelines Git-init" \
    io.k8s.display-name="Red Hat OpenShift Pipelines Git-init" \
    io.k8s.description="git-init is a binary that makes it easy to clone a repository from a Tekton Task. It is usually used via the git-clone Tasks." \  
    io.openshift.tags="pipelines,tekton"

RUN mkdir /licenses
COPY LICENSE /licenses/.
# Not adding the user to make sure git+ssh uses HOME to read client configuration
RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot -d /home/git -m nonroot
USER 65532

ENTRYPOINT ["/ko-app/git-init"]
