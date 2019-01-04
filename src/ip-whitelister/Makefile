GOOS=linux
GOARCH=amd64
GOCMD=go
GOBUILD=GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u
BINARY_NAME=ipwhitelister

all: deps test build package 
build:
	$(GOBUILD) -o $(BINARY_NAME) -v 
test:
	$(GOTEST) -v ./...
package:
clean:
	$(goclean)
	rm -f $(BINARY_NAME)
deps:
	# this is needed to be able to do lambda "stuff"
	$(GOGET) github.com/aws/aws-lambda-go/lambda
	$(GOGET) github.com/aws/aws-sdk-go
