.PHONY: generate
generate:
	tools/format.sh
	go run tools/generate/generate.go
	tools/generate.sh

.PHONY: init
init:
	git submodule update --init

.PHONY: update
update:
	git submodule update --remote
