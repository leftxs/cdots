GOCMD=go
BINARY_NAME=cdots
GOBUILD=$(GOCMD) build

all: run
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
run: build
	./$(BINARY_NAME)

