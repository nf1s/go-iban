build:
	@go build -gcflags="all=-N -l" -o iban

run:
	@go run .

pid proc:
	@ps aux | grep {{proc}}

debug pid:
	@arch -arm64 dlv attach {{pid}}

test:
	@go test
