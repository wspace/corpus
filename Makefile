.PHONY: all
all: format_projects README.md assembly.md building.md tidy_submodules

.PHONY: format_projects
format_projects: projects.json tools/format_projects.bash
	@echo 'Formatting projects.json'
	@tools/format_projects.bash

README.md: projects.json README.md.tmpl tools/generate/generate.go
	@echo 'Generating README.md'
	@go run tools/generate/generate.go

assembly.md: projects.json tools/generate_assembly.jq
	@echo 'Generating assembly.md'
	@jq -rf tools/generate_assembly.jq projects.json | tools/sponge assembly.md

building.md: projects.json tools/generate_building.jq
	@echo 'Generating building.md'
	@jq -rf tools/generate_building.jq projects.json | tools/sponge building.md

.PHONY: tidy_submodules
tidy_submodules: projects.json tools/add_submodules.jq tools/format_gitmodules.bash
	@echo 'Tidying Git submodules'
	@jq -rf tools/add_submodules.jq projects.json | sh
	@tools/format_gitmodules.bash
	@git add .gitmodules

.PHONY: init_submodules
init_submodules:
	git submodule update --init

.PHONY: update_submodules
update_submodules:
	git submodule update --remote
