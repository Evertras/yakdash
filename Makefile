.PHONY: default
default:
	go run ./cmd/yakdash/main.go

bin/yakdash:
	go build -o bin/yakdash ./cmd/yakdash
