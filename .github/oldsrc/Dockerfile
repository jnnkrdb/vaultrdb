# ----------------------------------------------- 
# Building the backend service and the operator with
# go build tools

# Build the manager binary
FROM golang:1.19 as builder

RUN mkdir /workspace
WORKDIR /workspace

# Copy the Go Modules manifests
COPY backend/go.mod go.mod
COPY backend/go.sum go.sum

# Copy the go source
COPY backend/main.go main.go
COPY backend/api/ api/
COPY backend/controllers/ controllers/
COPY backend/svc/ svc/

ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /vaultrdb .


# ----------------------------------------------- 
# Building the frontend with angular, which the backend 
# uses to host the frontend


# ----------------------------------------------- 
# Aggregating all compiled binaries and files into the alpine
# base container with the correct
FROM alpine:3.10

# Create the neccessary directories
# and set the Workdir
RUN mkdir -p /usr/share/vaultrdb/{swaggerui,frontend}
WORKDIR /usr/share/vaultrdb

# Copy swaggerui files
COPY swagger.yaml /usr/share/vaultrdb/swaggerui/swagger.yaml

# Copy frontend files

# Copy Version into the directory
COPY VERSION ./VERSION

# Copy the controller binary to the bin location
# and set the correct entrypoint
COPY --from=builder /vaultrdb /usr/local/bin/vaultrdb
RUN chmod a+x /usr/local/bin/vaultrdb
ENTRYPOINT ["/usr/local/bin/vaultrdb"]