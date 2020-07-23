.PHONY: run
run:
	go run -v ./cmd/apiserver

.DEFAULT_GOAL := run