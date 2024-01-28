# ----------------------------------------------- 
# Building the operator the go build tools and the crud api
FROM golang:1.19 as builder
WORKDIR /workspace

# copy the code files
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY startup.go startup.go
COPY api/ api/
COPY controllers/ controllers/
COPY crud/ crud/

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
FROM gcr.io/distroless/static:nonroot
WORKDIR /
# Copy Versions Information
RUN mkdir -p /vaultrdb/{swagger,ui,data}
COPY VERSION /vaultrdb/VERSION
# Copy Operators Binary
COPY --from=builder /vaultrdb /usr/local/bin/vaultrdb
RUN chmod a+x /usr/local/bin/vaultrdb
RUN chmod a+x -R /vaultrdb
USER 65532:65532
RUN chown 65532:65532 /usr/local/bin/vaultrdb
RUN chmod 65532:65532 -R /vaultrdb
ENTRYPOINT ["vaultrdb"]
