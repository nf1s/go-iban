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
