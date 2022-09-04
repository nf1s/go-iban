build:
	@go build .

run:
	@go run .

debug:
	@arch -arm64 dlv debug main.go

test:
	@go test
