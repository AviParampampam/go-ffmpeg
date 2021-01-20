.PHONY: build
build:
	go build cmd/metis/main.go

.DEFAULT_GOAL := build