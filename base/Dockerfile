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
RUN mkdir -p /opt/vaultrdb &&\
    mkdir -p /opt/vaultrdb/config &&\
    mkdir -p /opt/vaultrdb/web &&\
    mkdir -p /opt/vaultrdb/swagger &&\
    mkdir -p /opt/vaultrdb/entrypoint.d &&\
    mkdir -p /opt/vaultrdb/temp &&\
    mkdir -p /opt/vaultrdb/data
# Copy the VaultRDB Directory
COPY vaultrdb/config/ /opt/vaultrdb/config/
COPY vaultrdb/entrypoint.d/ /opt/vaultrdb/entrypoint.d/
# Set the Entrypoint Script
ENTRYPOINT ["/opt/vaultrdb/entrypoint.sh"]