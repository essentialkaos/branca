################################################################################

.DEFAULT_GOAL := help
.PHONY = fmt deps deps-test test help

################################################################################

deps: ## Download dependencies
	go get -d -v golang.org/x/crypto/chacha20poly1305

deps-test: ## Download dependencies for tests
	git config --global http.https://pkg.re.followRedirects true
	go get -d -v pkg.re/check.v1

test: ## Run tests
	go test -covermode=count .

bench: ## Run benchmarks
	go test -check.b

gen-fuzz: ## Generate go-fuzz archives
	go-fuzz-build -func FuzzEncode -o "branca-enc-fuzz.zip" github.com/essentialkaos/branca
	go-fuzz-build -func FuzzDecode -o "branca-dec-fuzz.zip" github.com/essentialkaos/branca

fmt: ## Format source code with gofmt
	find . -name "*.go" -exec gofmt -s -w {} \;

help: ## Show this info
	@echo -e '\nSupported targets:\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-12s\033[0m %s\n", $$1, $$2}'
	@echo -e ''

################################################################################
