.PHONY: test
test:
	go clean -testcache && go test -v -cover ./game

.PHONY: build
build:
	go build -o ./bin/gol .

.PHONY: run
run:
	go run . run

.PHONY: help
help:
	go run . -h

.PHONY: check-releaser
check-releaser:
	goreleaser check
	goreleaser release --snapshot --clean
