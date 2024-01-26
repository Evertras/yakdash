.PHONY: default
default:
	go run ./cmd/yakdash/main.go

bin/yakdash:
	go build -o bin/yakdash ./cmd/yakdash

.PHONY: fmt
fmt: node_modules
	go fmt ./...
	npx prettier --write .

node_modules: package.json package-lock.json
	@npm install
	@touch node_modules
