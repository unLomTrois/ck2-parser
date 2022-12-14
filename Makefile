cmd = ./cmd/template/
output = ./bin/

all = build run
.PHONY: all

build:
	mkdir -p $(output)
	go build -v -o $(output) $(cmd)

run:
	go run $(cmd)

.DEFAULT_GOAL := build
