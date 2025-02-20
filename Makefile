.DEFAULT_GOAL := help
.PHONY: gen fmt lint help

gen: ## コードを生成する
	@go tool ogen -clean -package zondaapi -target ./lib/go/zondaapi ./api/openapi.yaml

fmt: ## コードを整形する
	@go tool goimports -w .

lint: ## コードの静的解析を実装する
	@go vet $$(go list ./... | grep -v -e /zondaapi)
	@go tool staticcheck $$(go list ./... | grep -v -e /zondaapi)

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) \
      | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
