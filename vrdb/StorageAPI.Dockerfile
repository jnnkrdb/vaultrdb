# ----------------------------------------------- 
# Building the go binary
FROM golang:1.21-alpine as builder
RUN apk add --no-cache --update gcc g++
WORKDIR /workspace/src.go
# copy the code files
COPY vrdb.go/ /vrdb.go/
COPY storageapi/src.go/ /workspace/src.go/
# set env vars
ENV CGO_ENABLED=1
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /vaultrdb-storageapi .
# ----------------------------------------------- 
# Finish the operator, api, ui, storage build with the final image
FROM vaultrdbbase:latest as final
# Copy the VaultRDB Directory Contents
COPY storageapi/vaultrdb/entrypoint.d/ /opt/vaultrdb/entrypoint.d/
# Copy Operators Binary
COPY --from=builder /vaultrdb-storageapi /usr/local/bin/vaultrdb-storageapi
# Set the user for the config and the operator binaries
RUN chmod a+x /usr/local/bin/vaultrdb-storageapi &&\
    chmod a+x -R /opt/vaultrdb &&\
    chown 65532:65532 /usr/local/bin/vaultrdb-storageapi &&\
    chown 65532:65532 -R /opt/vaultrdb
# set the entrypoints
USER 65532:65532
EXPOSE 8080
# set the entrypoints
CMD [ "vaultrdb-storageapi" ]
