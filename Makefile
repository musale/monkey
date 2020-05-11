.PHONY: coverprofile run tests
run:
	@go run cmd/*

tests:
	@go test -cover ./...

coverprofile:
	@go test -coverprofile /tmp/profile.out ./...
	@go tool cover -html=/tmp/profile.out