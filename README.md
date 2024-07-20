# IBAN Validator

A web service written in go that validates IBAN numbers.

## Requirements

- Go 1.19
- Postgres 14
- Just

## Installation

```bash
$ go mod download
```

## Development

Make sure your Postgres DB is running and you have the following environment variables set:

```bash
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=

```

Run the migrations script to create the necessary tables:

```bash

just migrate

```

Run the server:

```bash

just run

```

Build the project:

```bash
just Build
```

## Deployment

### docker-compose

```bash
$ docker-compose up
```

### Kubernetes

```bash

$ just deploy

```
