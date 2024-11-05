# ----------------------------------------------- 
# Building the go binary
FROM golang:1.19.13 AS operator
WORKDIR /workspace/src.go

# copy the code files
COPY ops/ /workspace/src.go

# set env vars
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux

# START BUILD
RUN go mod download && go build -o /vaultrdb-operator .

# ----------------------------------------------- 
# Building the go binary
FROM node:latest AS frontend
WORKDIR /workspace/src.angular
# copy the code files
COPY ui/ /workspace/src.angular

RUN npm install 

RUN npm run build

# ----------------------------------------------- 
# Finish the operator, api, ui build with the final image
# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# FROM gcr.io/distroless/static:nonroot
FROM alpine:3.19
WORKDIR /

# install neccessary binaries
RUN apk add --no-cache --update openssl

# Copy the VaultRDB Directory Contents
COPY vaultrdb/ /opt/vaultrdb

# Copy Operators Binary and Frontend Files
COPY --from=operator /vaultrdb-operator /usr/local/bin/vaultrdb-operator
COPY --from=frontend /workspace/src.angular/dist/vaultrdb/ /opt/vaultrdb/web

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