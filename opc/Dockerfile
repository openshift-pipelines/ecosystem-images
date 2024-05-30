FROM registry.access.redhat.com/ubi9/go-toolset:1.21.9-1.1715774364@sha256:f001ad1001a22fe5f6fc7d876fc172b01c1b7dcd6c498f83a07b425e24275a79 AS builder

WORKDIR /app
COPY . .
RUN go build -buildvcs=false -mod=vendor -o /app/opc main.go

FROM registry.access.redhat.com/ubi9/ubi-minimal:latest@sha256:ef6fb6b3b38ef6c85daebeabebc7ff3151b9dd1500056e6abc9c3295e4b78a51
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
