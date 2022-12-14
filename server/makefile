task1:
	@echo "task1"

task2: task1
	@echo "task2"

task3: task1 task2
	@echo "task3"

test:
	go test -v

BUILD_TAGS_PRODUCTION='production'
BUILD_TAGS_DEVELOPMENT='development unittest'

build-base:
	go build -o ./build/${BIN_NAME} -tags '$(BUILD_TAGS) go' -installsuffix go -ldflags '-s -w' main.go

build-development:
	$(MAKE) build-base BUILD_TAGS=${BUILD_TAGS_DEVELOPMENT} BIN_NAME=dev

build-production-windows:
	$(MAKE) build-base BUILD_TAGS=${BUILD_TAGS_PRODUCTION} CGO_ENABLED=0 GOOS=windows BIN_NAME=prod

build-production-linux:
	$(MAKE) build-base BUILD_TAGS=${BUILD_TAGS_PRODUCTION} CGO_ENABLED=0 GOOS=linux GOARCH=amd64 BIN_NAME=prod

