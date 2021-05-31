# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOBENCH=$(GOTEST) -bench=. -benchmem
BINARY_NAME=merger
BINARY_UNIX=$(BINARY_NAME)_unix
# no deps, no distribution, no go mod needed
GO111MODULE=off

all: clean test bench build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
test:
		$(GOTEST) -v ./...
bench:
		$(GOBENCH)
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
