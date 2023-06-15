# Build the manager binary
FROM golang:1.19 as builder

RUN mkdir /workspace
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY svc/ svc/

ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /vaultrdb .

FROM alpine:3.10
WORKDIR /
COPY --from=builder /vaultrdb /vaultrdb
RUN chmod a+x /vaultrdb
ENTRYPOINT ["/vaultrdb"]