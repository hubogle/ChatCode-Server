.DEFAULT_GOAL := help
GIT_TAG = $(shell git rev-parse --short=8 HEAD)

## run: 运行服务
.PHONY: run
run:
	go run ./cmd/server/main.go

## gen_model: 生成 model 文件
.PHONY: gen_model
gen_model:
	go run ./cmd/generate/main.go

## help: 帮助信息
.PHONY: help
help: Makefile
	@echo "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
