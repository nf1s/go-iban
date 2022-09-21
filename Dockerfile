# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o iban

## Deploy
FROM golang:1.19-buster

WORKDIR /build
COPY --from=build /app/iban /build/iban

EXPOSE 8080
