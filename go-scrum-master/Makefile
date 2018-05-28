GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=sm
BINARY_ZIP=scrummaster.zip


build:
	GOARCH=amd64 GOOS=linux	$(GOBUILD) -o $(BINARY_NAME) -v

zip: clean build
	zip $(BINARY_ZIP) $(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_ZIP)
