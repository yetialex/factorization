BUILD_FOLDER := dist

.PHONY: all clean build

all: test build

build: ## Build binary and docker images
	docker build -t build.factorization --force-rm -f build.Dockerfile .
	mkdir -p $(CURDIR)/$(BUILD_FOLDER)
	docker run -v $(CURDIR)/$(BUILD_FOLDER):/opt/mount --rm --entrypoint cp build.factorization /src/app/dist/prime_factorization /opt/mount/prime_factorization
	docker build -t factorization --force-rm -f Dockerfile .

test: ## Start unit-tests
	go test -v ./...

clean: ## Remove artifacts
	rm -rf $(BUILD_FOLDER)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
