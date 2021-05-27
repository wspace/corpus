.PHONY: all
all: format_projects licenses README.md assembly.md building.md tidy_submodules

.PHONY: format_projects
format_projects: projects.json tools/format_projects.sh
	@echo 'Formatting projects.json'
	@tools/format_projects.sh

.PHONY: licenses
licenses: projects.json tools/licenses/licenses.go
	@echo 'Getting licenses'
	@go run tools/licenses/licenses.go | underscore print | tools/sponge projects.json

README.md: projects.json README.md.tmpl tools/generate/generate.go
	@echo 'Generating README.md'
	@go run tools/generate/generate.go

assembly.md: projects.json tools/generate_assembly.jq
	@echo 'Generating assembly.md'
	@jq -rf tools/generate_assembly.jq projects.json > assembly.md

building.md: projects.json tools/generate_building.jq
	@echo 'Generating building.md'
	@jq -rf tools/generate_building.jq projects.json > building.md

.PHONY: tidy_submodules
tidy_submodules: projects.json tools/add_submodules.jq tools/format_gitmodules.sh
	@echo 'Tidying Git submodules'
	@jq -rf tools/add_submodules.jq projects.json | sh
	@tools/format_gitmodules.sh
	@git add .gitmodules

.PHONY: init_submodules
init_submodules:
	git submodule update --init

.PHONY: update_submodules
update_submodules:
	git submodule update --remote
