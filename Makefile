
VERSION=$(shell git describe --tags)
GO_LDFLAGS=$(shell linkflags -pkg=.)

all: clean lint build package

clean:
	rm -rf out/*

lint:
	go vet .
	golangci-lint run

build:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_LDFLAGS)" -o out/linux-amd64/clico
	GOOS=linux GOARCH=arm64 go build -ldflags "$(GO_LDFLAGS)" -o out/linux-arm64/clico
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(GO_LDFLAGS)" -o out/macos-amd64/clico
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(GO_LDFLAGS)" -o out/macos-arm64/clico
	GOOS=windows GOARCH=amd64 go build -ldflags "$(GO_LDFLAGS)" -o out/windows-amd64/clico.exe
	GOOS=windows GOARCH=arm64 go build -ldflags "$(GO_LDFLAGS)" -o out/windows-arm64/clico.exe

package:
	tar -czvf out/clico-$(VERSION)-linux-arm64.tar.gz out/linux-arm64/clico
	tar -czvf out/clico-$(VERSION)-linux-amd64.tar.gz out/linux-amd64/clico
	tar -czvf out/clico-$(VERSION)-macos-amd64.tar.gz out/macos-amd64/clico
	tar -czvf out/clico-$(VERSION)-macos-arm64.tar.gz out/macos-arm64/clico
	tar -czvf out/clico-$(VERSION)-windows-amd64.zip out/windows-amd64/clico.exe
	tar -czvf out/clico-$(VERSION)-windows-arm64.zip out/windows-arm64/clico.exe

tools:
	go install github.com/gravitational/version/cmd/linkflags
