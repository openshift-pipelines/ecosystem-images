FROM quay.io/projectquay/golang:1.22 AS builder
# FROM registry.access.redhat.com/ubi9/go-toolset:1.21.10-1.1719562237@sha256:ae17d73e70a966f39ef4dfca74241e3ca4374cd1198b02c30ea0748b8dcc83a6 AS builder

WORKDIR /app
COPY opc .
RUN go build -buildvcs=false -mod=vendor -o /app/opc main.go

FROM registry.access.redhat.com/ubi9/ubi-minimal:latest@sha256:104cf11d890aeb7dd5728b7d7732e175a0e4018f1bb00d2faebcc8f6bf29bd52
LABEL name="opc" \
      com.redhat.component="opc" \
      io.k8s.display-name="opc" \
      summary="A CLI for OpenShift Pipeline" \
      description="opc makes it easy to work with Tekton resources in OpenShift Pipelines. It is built on top of tkn and tkn-pac and expands their capablities to the functionality and user-experience that is available on OpenShift." \
      io.k8s.description="opc makes it easy to work with Tekton resources in OpenShift Pipelines. It is built on top of tkn and tkn-pac and expands their capablities to the functionality and user-experience that is available on OpenShift." \
      io.openshift.tags="pipelines,tekton"
COPY --from=builder /app/opc /usr/bin
RUN mkdir /licenses
COPY LICENSE /licenses/.
RUN microdnf install -y shadow-utils
RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532
