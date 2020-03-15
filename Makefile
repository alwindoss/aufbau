
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
BINARY_NAME=aufbau
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build: 
	$(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)/...
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_LOC)
run:
	$(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)/...
	./$(BINARY_LOC)/$(BINARY_NAME)


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LOC)/$(BINARY_UNIX) -v ./cmd/$(BINARY_NAME)/...
docker-build:
	docker build -t latest .