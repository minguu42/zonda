.DEFAULT_GOAL := help
.PHONY: setup gen fmt lint help

setup: ## 開発に必要なツールをインストールする
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/ogen-go/ogen/cmd/ogen@latest

gen: ## コードを生成する
	@ogen -clean -package oapi -target ./api/oapi ./api/openapi.yaml

fmt: ## コードを整形する
	@goimports -w .

lint: ## コードの静的解析を実装する
	@go vet $$(go list ./... | grep -v -e /oapi)
	@staticcheck $$(go list ./... | grep -v -e /oapi)

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) \
      | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
