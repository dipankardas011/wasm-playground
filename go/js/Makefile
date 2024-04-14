
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\n  \033[36mHELLO FROM BUILDER\033[0m\n\n\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ Build
.PHONY: build
build: ## build the main.wasm
	GOOS=js GOARCH=wasm go build -o main.wasm .


##@ Run
.PHONY: run
run: ## run the main.wasm to port 8080
	# go install github.com/shurcooL/goexec
	# goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
	go run http.go
