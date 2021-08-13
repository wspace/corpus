PROJECTS = $(wildcard */*.json)

.PHONY: all
all: licenses README.md assembly.md building.md tidy_submodules

# TODO fix
.PHONY: format
format: projects.json tools/format_projects.sh
	@echo 'Formatting projects'
	@tools/format_projects.sh

.PHONY: licenses
licenses: $(PROJECTS) tools/licenses/licenses.go
	@echo 'Getting licenses'
	@go run tools/licenses/licenses.go

README.md: $(PROJECTS) README.md.tmpl tools/generate.go tools/generate/generate.go
	@echo 'Generating README.md'
	@go run tools/generate/generate.go

assembly.md: $(PROJECTS) tools/generate_assembly.jq
	@echo 'Generating assembly.md'
	@jq -rsf tools/generate_assembly.jq $(PROJECTS) > assembly.md

building.md: $(PROJECTS) tools/generate_building.jq
	@echo 'Generating building.md'
	@jq -rsf tools/generate_building.jq $(PROJECTS) > building.md

.PHONY: tidy_submodules
tidy_submodules: $(PROJECTS) tools/submodules/submodules.go tools/format_gitmodules.sh
	@echo 'Tidying Git submodules'
	@go run tools/submodules/submodules.go
	@tools/format_gitmodules.sh
	@git add .gitmodules

# Clone all submodules
.PHONY: init_submodules
init_submodules:
	git submodule update --init

# Update all submodules to latest remote head
.PHONY: update_submodules
update_submodules:
	git submodule update --remote
