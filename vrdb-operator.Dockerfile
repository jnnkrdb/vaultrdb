# ---------------------------------------------------------------------------------------------- Golang
# Building the go binary
#FROM golang:1.19 AS operator
FROM golang:1.23 AS operator
WORKDIR /workspace

# copy the code files
COPY pkg/ /workspace/pkg
COPY bin/ /workspace/bin
COPY go.mod /workspace/go.mod
COPY go.sum /workspace/go.sum

# set env vars
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux

# START BUILD
RUN go mod download && go build -o /vaultrdb-operator ./bin/vaultrdb-operator/main.go

# ---------------------------------------------------------------------------------------------- Final Alpine
FROM alpine:3.19
WORKDIR /

# install neccessary binaries
RUN apk add --no-cache --update openssl

# Copy the VaultRDB Directory Contents
COPY vaultrdb/ /opt/vaultrdb

# Copy Operators Binary and Frontend Files
COPY --from=operator /vaultrdb-operator /usr/local/bin/vaultrdb-operator

# Set the user for the config and the operator binaries
RUN chmod a+x /usr/local/bin/vaultrdb-operator &&\
    chmod a+x -R /opt/vaultrdb &&\
    chown 65532:65532 /usr/local/bin/vaultrdb-operator &&\
    chown 65532:65532 -R /opt/vaultrdb
    
USER 65532:65532

# set the entrypoints
EXPOSE 80
ENTRYPOINT ["/opt/vaultrdb/entrypoint.sh"]
CMD [ "vaultrdb-operator" ]