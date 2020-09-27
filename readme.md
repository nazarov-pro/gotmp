# Go

## Go Lang Test

* Go Test all packages(with coverage): `go test ./... -cover`
* Go Test all packages(export test output): `go test ./... -cover -coverprofile=bin/c.out`
* Go generate html from output file: `go tool cover -html=bin/c.out -o bin/coverage.html`
* Go clean test cache: `go clean -testcache`
