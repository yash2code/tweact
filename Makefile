.SUFFIXES:

SOURCE := $(shell find pkg cmd protoc-gen-tweact -name '*.go')

build: bin/api
	@: # shhh
.PHONY: build

dev:
	@CompileDaemon \
		-exclude-dir web \
		-build make \
		-command bin/api
.PHONY: dev

test:
	@go test ./...
.PHONY: test

generate: vendor _tools/bin/protoc-gen-tweact
	retool do protoc -I pkg:rpc:vendor \
		--doc_out=html,doc/index.html \
		--lint_out=. \
		--go_out=pkg \
		--twirp_out=pkg \
		--tweact_out=web/rpc \
		rpc/todos/v1/service.proto rpc/demo/service.proto
.PHONY: generate

vendor:
	dep ensure
.PHONY: vendor

_tools/bin/%: $(SOURCE)
	go build -o $@ $*/main.go

bin/api: $(SOURCE)
	go build -o $@ cmd/$*/$*.go

bin/web:
	cd web && yarn build && yarn pkg