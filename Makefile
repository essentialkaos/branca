################################################################################

# This Makefile generated by GoMakeGen 1.3.0 using next command:
# gomakegen .
#
# More info: https://kaos.sh/gomakegen

################################################################################

.DEFAULT_GOAL := help
.PHONY = fmt vet git-config deps deps-test test gen-fuzz help

################################################################################

git-config: ## Configure git redirects for stable import path services
	git config --global http.https://pkg.re.followRedirects true

deps: git-config ## Download dependencies
	go get -d -v golang.org/x/crypto/chacha20poly1305

deps-test: git-config ## Download dependencies for tests
	go get -d -v pkg.re/check.v1

test: ## Run tests
	go test -covermode=count .

bench: ## Run benchmarks
	go test -check.b

gen-fuzz: ## Generate archives for fuzz testing
	which go-fuzz-build &>/dev/null || go get -u -v github.com/dvyukov/go-fuzz/go-fuzz-build
	go-fuzz-build -o fuzz.zip github.com/essentialkaos/branca

fmt: ## Format source code with gofmt
	find . -name "*.go" -exec gofmt -s -w {} \;

vet: ## Runs go vet over sources
	go vet -composites=false -printfuncs=LPrintf,TLPrintf,TPrintf,log.Debug,log.Info,log.Warn,log.Error,log.Critical,log.Print ./...

help: ## Show this info
	@echo -e '\n\033[1mSupported targets:\033[0m\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-12s\033[0m %s\n", $$1, $$2}'
	@echo -e ''
	@echo -e '\033[90mGenerated by GoMakeGen 1.3.0\033[0m\n'

################################################################################
