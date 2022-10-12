url := 'https://app.circleci.com/pipelines/github/nf1s/iban?branch=master'

default:
	@just --list

build:
	@go build -gcflags="all=-N -l" -o iban

run:
	@go run .

pid proc:
	@ps aux | grep {{proc}}

attach pid:
	@arch -arm64 dlv attach {{pid}}

debug:
	@arch -arm64 dlv debug .

debug-tests:
	@arch -arm64 dlv test .

test:
	@go test

cli arg:
	@./cli/cli {{arg}}

view:
  @{{ if os() =~ "macos.*" { "open " +url } else { "xdg-open "+url } }}

migrate:
	@./migrate -source file://migrations -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable up

deploy:
	@skaffold run

delete:
	@skaffold delete

manifests:
	@./scripts/manifests.sh
