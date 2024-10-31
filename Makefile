BINARY_DIR=bin
VERSION?=1.0.0
BUILD_TIME=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

.PHONY: all clean build test

all: build

build:
	go build ${LDFLAGS} -o ${BINARY_DIR}/simulator ./cmd/simulator