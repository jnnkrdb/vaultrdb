# ----------------------------------------------- GO-BUILD
# BUILDING STAGE
FROM golang:1.20 AS build-go
# CREATE BUILD-DIR
RUN mkdir /build
WORKDIR /build
# COPY APPLICATION
COPY go/ /build/
# SET ENVs for GO BUILD
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /gobinary
# -----------------------------------------------
# FINAL STAGE
FROM alpine:3.16
# GET gobinary
RUN mkdir -p /app/ui
COPY --from=build-go /gobinary /app/gobinary
COPY ui/ /app/ui/
RUN chmod a+x /app/gobinary
# Define ENTRYPOINT
ENTRYPOINT ["/app/gobinary"]