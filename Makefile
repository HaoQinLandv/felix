LDFLAGS := "-s -w -X main.buildTime=$(shell date -u '+%Y-%m-%dT%I:%M:%S%p') -X main.gitHash=$(shell git rev-parse HEAD)"

run: build
	./build/felix -V

build:
	go build -race -ldflags $(LDFLAGS)  -o build/felix *.go

release:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags $(LDFLAGS) -o release/felix-amd64-darwin *.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -o release/felix-amd64-win.exe *.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o release/felix-amd64-linux *.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags $(LDFLAGS) -o release/felix-amd64-linux-arm *.go

.PHONY: release

