# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o iban
RUN cd cli && go build -o cli

## Deploy
FROM golang:1.19-buster

WORKDIR /build
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
COPY --from=build /app/iban /build/iban
COPY --from=build /app/cli/cli /build/cli/cli
COPY --from=build /app/migrations/ /build/migrations
COPY --from=build /app/data/ /build/data

EXPOSE 8080
