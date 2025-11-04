VERSION := $(shell git describe --tags)
COMMIT  := $(shell git log -1 --format='%H')

all: build

LD_FLAGS = -X ulst-relay/cmd.Version=$(VERSION) \
	-X ulst-relay/cmd.Commit=$(COMMIT) \

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

get:
	@echo "  >  \033[32mDownloading & Installing all the modules...\033[0m "
	go mod tidy && go mod download

build:
	@echo " > \033[32mBuilding ulst-relay...\033[0m "
	go build -mod readonly $(BUILD_FLAGS) -o build/ulst-relay main.go

install: build
	sudo mv build/ulst-relay /usr/local/bin/

build-linux:
	@GOOS=linux GOARCH=amd64 go build --mod readonly $(BUILD_FLAGS) -o ./build/ulst-relay main.go

abi:
	@echo " > \033[32mGenabi...\033[0m "
	abigen --abi ./bindings/stake_manager/StakeManager_abi.json --pkg stake_manager --type StakeManager --out ./bindings/stake_manager/StakeManager.go
	abigen --abi ./bindings/stake_pool/StakePool_abi.json --pkg stake_pool --type StakePool --out ./bindings/stake_pool/StakePool.go

clean:
	@echo " > \033[32mCleaning build files ...\033[0m "
	rm -rf build
fmt :
	@echo " > \033[32mFormatting go files ...\033[0m "
	go fmt ./...

get-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest

lint:
	golangci-lint run ./... --skip-files ".+_test.go"

.PHONY: all lint clean build
