# ----------------------------------------------- 
# Building the operator the go build tools
FROM golang:1.19 as builder
WORKDIR /workspace
# copy the code files
COPY src/ /workspace/
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /vaultrdb .

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
# Copy Versions Information
RUN mkdir -p /usr/share/vaultrdb
COPY VERSION /usr/share/vaultrdb/VERSION
# Copy Operators Binary
COPY --from=builder /vaultrdb /usr/local/bin/vaultrdb
RUN chmod a+x /usr/local/bin/vaultrdb
USER 65532:65532
ENTRYPOINT ["/usr/local/bin/vaultrdb"]
