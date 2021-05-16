.PHONY: generate
all: format README.md assembly.md building.md

.PHONY: format
format: tools/format.bash
	tools/format.bash

README.md: projects.json README.md.tmpl tools/generate/generate.go
	go run tools/generate/generate.go

assembly.md: projects.json tools/generate_assembly.jq
	jq -rf tools/generate_assembly.jq projects.json > assembly.md

building.md: projects.json tools/generate_building.jq
	jq -rf tools/generate_building.jq projects.json > building.md

.PHONY: init
init:
	git submodule update --init

.PHONY: update
update:
	git submodule update --remote
