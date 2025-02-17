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
RUN go mod download && go build -o /vaultrdb ./bin/vaultrdb/main.go

# ---------------------------------------------------------------------------------------------- Frontend
# Building the go binary
FROM node:22.11 AS frontend
WORKDIR /workspace/frontend
# copy the code files
COPY ui/vue/ /workspace/frontend
RUN npm install 
RUN npm run build

# ---------------------------------------------------------------------------------------------- Final Alpine
FROM alpine:3.19
WORKDIR /

# install neccessary binaries
RUN apk add --no-cache --update openssl

# Copy the VaultRDB Directory Contents
COPY vaultrdb/ /opt/vaultrdb

# create vault user with home dir
RUN useradd -r -d /opt/vaultrdb/home -s /bin/sh -g vault -u 3454 vault

# Copy Operators Binary and Frontend Files
COPY --from=operator /vaultrdb /usr/local/bin/vaultrdb
COPY --from=frontend /workspace/frontend/dist/ /opt/vaultrdb/web

# Set the user for the config and the operator binaries
#RUN chmod a+x /usr/local/bin/vaultrdb &&\
#    chmod a+x -R /opt/vaultrdb &&\
RUN chown vault:vault /usr/local/bin/vaultrdb &&\
    chown vault:vault -R /opt/vaultrdb
    
USER vault:vault

# set the entrypoints
ENTRYPOINT ["/opt/vaultrdb/entrypoint.sh"]
CMD [ "vaultrdb" ]