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

test:
	@go test
