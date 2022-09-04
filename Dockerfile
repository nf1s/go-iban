# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /build

## Deploy
FROM golang:1.18-buster

COPY --from=build /build /build

EXPOSE 8080


ENTRYPOINT ["/build"]
