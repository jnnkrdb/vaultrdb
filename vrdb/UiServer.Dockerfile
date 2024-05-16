# ----------------------------------------------- 
# Building the go binary
FROM golang:1.21.1 as builder
WORKDIR /workspace/src.go
# copy the code files
COPY vrdb.go/ /vrdb.go/
COPY uiserver/src.go/ /workspace/src.go/
# set env vars
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /vaultrdb-ui .
# ----------------------------------------------- 
# Finish the operator, api, ui, storage build with the final image
FROM vaultrdbbase:latest as final
# Copy the VaultRDB Directory Contents
COPY uiserver/vaultrdb/entrypoint.d/ /opt/vaultrdb/entrypoint.d/
COPY uiserver/vaultrdb/swagger/ /opt/vaultrdb/swagger/
# copy html contents from frontend build
COPY uiserver/src.html/ /opt/vaultrdb/web/
# Copy Operators Binary
COPY --from=builder /vaultrdb-ui /usr/local/bin/vaultrdb-ui
# Set the user for the config and the operator binaries
RUN chmod a+x /usr/local/bin/vaultrdb-ui &&\
    chmod a+x -R /opt/vaultrdb &&\
    chown 65532:65532 /usr/local/bin/vaultrdb-ui &&\
    chown 65532:65532 -R /opt/vaultrdb
# set the entrypoints
USER 65532:65532
EXPOSE 80
# set the entrypoints
CMD [ "vaultrdb-ui" ]
