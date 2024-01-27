GO_FILES=$(shell find pkg -name '*.go')

################################################################################
# Demo

.PHONY: default
demo: git-hooks
	@go run ./cmd/yakdash/main.go -c examples/clocks.yaml

################################################################################
# Build
bin/yakdash: git-hooks $(GO_FILES)
	go build -o bin/yakdash ./cmd/yakdash

################################################################################
# Test

.PHONY: test
test: git-hooks
	@go test ./...

.PHONY: test-coverage
test-coverage: coverage.out
	@go tool cover -html=coverage.out

coverage.out: $(GO_FILES)
	@go test -coverprofile=coverage.out ./pkg/...

################################################################################
# Lint / Format
.PHONY: lint
lint: git-hooks
	golangci-lint run ./...

.PHONY: fmt
fmt: node_modules git-hooks
	go fmt ./...
	npx prettier --write .

node_modules: package.json package-lock.json
	@npm install
	@touch node_modules

.PHONY:
git-hooks: .git/hooks/pre-commit

.git/hooks/pre-commit: ./.evertras/pre-commit.sh
	@cp ./.evertras/pre-commit.sh ./.git/hooks/pre-commit
	@chmod +x ./.git/hooks/pre-commit
