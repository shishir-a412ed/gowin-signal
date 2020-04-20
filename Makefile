BINARY ?= gowin-signal.exe
export GO111MODULE=on
export GOOS=windows

default: build

.PHONY: build
build: main.go
	go build -o $(BINARY) .

.PHONY: clean
clean:
	rm -f $(BINARY)
