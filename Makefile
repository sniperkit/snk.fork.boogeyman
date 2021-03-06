# application meta info
NAME := boogeyman
VERSION= 1.2.6
REVISION := $(shell git rev-parse --short HEAD)
BUILDDATE := $(shell date '+%Y/%m/%d %H:%M:%S %Z')
GOVERSION := $(shell go version)
LDFLAGS := -X 'main.revision=$(REVISION)' -X 'main.version=$(VERSION)' -X 'main.buildDate=$(BUILDDATE)' -X 'main.goVersion=$(GOVERSION)'
ENTRYPOINT := main.go

all: dep production

dep:
	dep ensure

# production mode: make [production | pro | p]
production pro p: gobindata-production build-production test-production

# development mode: make [development | develop | dev | d]
development develop dev d: dep gobindata-development build-development

# buid
build-%:
	GOOS=linux GOARCH=amd64	go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS) -X 'main.mode=$*'" -o bin/$(NAME)-linux-64 ./$(ENTRYPOINT)
	GOOS=darwin GOARCH=amd64 go build -tags="$* netgo" -installsuffix netgo -ldflags "$(LDFLAGS) -X 'main.mode=$*'" -o bin/$(NAME)-darwin-64 ./$(ENTRYPOINT)

# test
test-%:
	go test -tags="$* netgo" ./...

update:
	dep ensure -update

# Generate Go file
gobindata-%:
	go-bindata -pkg config -o config/config_tml.go config.$*.tml

# publicing
deploy:
	mv ./bin/* ./public/
