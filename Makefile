.PHONY: coverprofile run tests fmt
fmt:
	@go fmt ./...
run:
	@go run cmd/*

tests:
	@go test -cover ./...

coverprofile:
	@go test -coverprofile /tmp/profile.out ./...
	@go tool cover -html=/tmp/profile.out