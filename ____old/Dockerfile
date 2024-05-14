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
RUN mkdir -p /opt/vaultrdb
# Copy the VaultRDB Directory
COPY vaultrdb/ /opt/vaultrdb/
COPY LICENSE /opt/vaultrdb/config/LICENSE
# Copy Operators Binary
COPY --from=builder /vaultrdb /usr/local/bin/vaultrdb
# Set the user for the config and the operator binaries
RUN chmod a+x /usr/local/bin/vaultrdb &&\
    chmod a+x -R /opt/vaultrdb &&\
    chown 65532:65532 /usr/local/bin/vaultrdb &&\
    chown 65532:65532 -R /opt/vaultrdb
# set the entrypoints
USER 65532:65532
ENTRYPOINT ["/opt/vaultrdb/entrypoint.sh"]
CMD [ "vaultrdb" ]