.PHONY: test
test:
	go clean -testcache && go test -v -cover ./...

.PHONY: build
build:
	go build -o ./bin/gol .

.PHONY: run
run:
	go run . run

.PHONY: check-releaser
check-releaser:
	goreleaser check
	goreleaser release --snapshot --clean
