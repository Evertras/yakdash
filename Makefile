.PHONY: default
demo: git-hooks
	@go run ./cmd/yakdash/main.go -c examples/config.yaml

.PHONY: test
test: git-hooks
	@go test ./...

bin/yakdash: git-hooks
	go build -o bin/yakdash ./cmd/yakdash

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
