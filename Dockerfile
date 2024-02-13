# ----------------------------------------------- 
# Building the operator the go build tools and the crud api
FROM golang:1.19 as builder
WORKDIR /workspace
# copy the code files
COPY src/ /workspace/
# set env vars
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /vaultrdb .

# ----------------------------------------------- 
# Building the frontend ui

# ----------------------------------------------- 
# Finish the operator, api, ui build with the final image
# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# FROM gcr.io/distroless/static:nonroot
FROM alpine:3.19
WORKDIR /
# install neccessary binaries
RUN apk add openssl
# Create the main vaultrdb working direct
RUN mkdir -p /vaultrdb
# Copy the VaultRDB Directory
COPY vaultrdb/ /vaultrdb/
COPY LICENSE /vaultrdb/LICENSE
# create other needed libraries
RUN mkdir /vaultrdb/temp
# Copy Operators Binary
COPY --from=builder /vaultrdb /usr/local/bin/vaultrdb
RUN chmod a+x /usr/local/bin/vaultrdb
RUN chmod a+x -R /vaultrdb
# set the user and run the operator binaries
RUN chown 65532:65532 /usr/local/bin/vaultrdb
RUN chown 65532:65532 -R /vaultrdb
# configure default env variables
ENV SLEEP_BEFORE_SERVICE_START="0"
ENV VAULTRDB_SERVICENAME=""
ENV BASICAUTH_USER="vault"
ENV BASICAUTH_PASS="vault"
ENV ENABLE_SWAGGERUI="false"
ENV FQDN="vaultrdb.kubernetes.docker.internal"
ENV _TEMPDIR="/vaultrdb/temp"
ENV _DEBUG=""
# set the entrypoints
USER 65532:65532
ENTRYPOINT ["/vaultrdb/entrypoint.sh"]
CMD [ "vaultrdb" ]