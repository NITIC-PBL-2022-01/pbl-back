# vim: ft=make

set dotenv-load := true

run: fmt
	go run .

build:
	go build .

check:
	go build ./...

watch target:
	find . -name '*.go' | entr -c just {{target}}

update:
	go get -u -v && go mod tidy

test:
	go test ./...

# calling gofmt before others because golines does not show any
# useful error diagnostics. It shows just something like `Exit status: 101`.
fmt:
	go fmt ./...
	go run github.com/segmentio/golines -m 120 -w `find . -name '*.go'`
	go run golang.org/x/tools/cmd/goimports -w `find . -name '*.go'`
