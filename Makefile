PROJECTS = $(filter-out tools/% .% _%, $(wildcard */*.json))
SUBMODULES = $(PROJECTS:.json=)

.PHONY: all
all: tidy_submodules licenses README.md assembly.md building.md

.PHONY: format
format: $(PROJECTS) tools/format/format.go
	@echo 'Formatting projects'
	@go run tools/format/format.go $(PROJECTS)

.PHONY: licenses
licenses: $(PROJECTS) tools/licenses/licenses.go
	@echo 'Getting licenses'
	@go run tools/licenses/licenses.go $(PROJECTS)

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
	git submodule update --init --jobs 5

# Update all submodules to latest remote head
.PHONY: update_submodules
update_submodules:
	git submodule update --remote --jobs 5

# Update all manually-enumerated submodules to latest remote head
.PHONY: update_submodules_force
update_submodules_force:
	git submodule foreach 'git -C $toplevel submodule update --remote $name'

.PHONY: list_project_json
list_project_json:
	$(foreach project,$(PROJECTS),$(info $(project)))
	@:
