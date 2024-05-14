# ----------------------------------------------- 
# Building the go binary
FROM golang:1.21.1 as builder
WORKDIR /workspace
# copy the code files
COPY src/ /workspace/
# set env vars
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /vaultrdb-operator .
# ----------------------------------------------- 
# Finish the operator, api, ui, storage build with the final image
FROM vaultrdbbase:latest as final
# Copy the VaultRDB Directory Contents
COPY operator/vaultrdb/entrypoint.d/ /opt/vaultrdb/entrypoint.d/
# Copy Operators Binary
COPY --from=builder /vaultrdb-operator /usr/local/bin/vaultrdb-operator
# Set the user for the config and the operator binaries
RUN chmod a+x /usr/local/bin/vaultrdb-operator &&\
    chmod a+x -R /opt/vaultrdb &&\
    chown 65532:65532 /usr/local/bin/vaultrdb-operator &&\
    chown 65532:65532 -R /opt/vaultrdb
# set the entrypoints
USER 65532:65532
# set the entrypoints
CMD [ "vaultrdb-operator" ]