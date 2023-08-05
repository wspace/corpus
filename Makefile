PROJECTS = $(filter-out tools/% .% _%, $(wildcard */*/project.json))
DOCKERFILES = $(filter-out tools/% .% _%, $(wildcard */Dockerfile */*/Dockerfile))
RUST_FORMAT_CRATE = $(shell find tools/format -type f)
GO_TOOLS_PACKAGE = tools/generate.go tools/licenses.go

.PHONY: all
all: format format_submodules licenses README.md assembly.md building.md challenges.md missing_submodules.md

.PHONY: format
format: build_stamp

build_stamp: $(PROJECTS) $(RUST_FORMAT_CRATE) tools/format/format.go $(GO_TOOLS_PACKAGE)
	$(info Formatting projects)
	@cargo run -q --bin corpus-format $(filter $(PROJECTS),$?)
	@go run tools/format/format.go $(filter $(PROJECTS),$?)
	@touch build_stamp

.PHONY: licenses
licenses: $(PROJECTS) tools/licenses/licenses.go $(GO_TOOLS_PACKAGE)
	$(info Getting licenses)
	@go run tools/licenses/licenses.go $(PROJECTS)

README.md: $(PROJECTS) README.md.tmpl tools/generate/generate.go $(GO_TOOLS_PACKAGE)
	$(info Generating README.md)
	@go run tools/generate/generate.go

assembly.md: $(PROJECTS) tools/generate_assembly.jq
	$(info Generating assembly.md)
	@jq -rsf tools/generate_assembly.jq $(PROJECTS) > assembly.md

building.md: $(PROJECTS) $(DOCKERFILES) tools/generate_building.jq
	$(info Generating building.md)
	@jq -rsf --arg dockerfiles "$(DOCKERFILES)" tools/generate_building.jq $(PROJECTS) > building.md

challenges.md: $(PROJECTS) tools/generate_challenges.jq
	$(info Generating challenges.md)
	@jq -rsf tools/generate_challenges.jq $(PROJECTS) > challenges.md

missing_submodules.md: $(PROJECTS) tools/generate_missing_submodules.jq
	$(info Generating missing_submodules.md)
	@jq -rsf tools/generate_missing_submodules.jq $(PROJECTS) > missing_submodules.md

.PHONY: docker_build
docker_build: tools/generate_docker_build.sh
	@tools/generate_docker_build.sh

.PHONY: format_submodules
format_gitmodules: $(PROJECTS)
	$(info Formatting Git submodules)
	@tools/format_gitmodules.sh
	@git add .gitmodules

# Clone all submodules
.PHONY: init_submodules
init_submodules:
# Ignore submodule worktree changes
	git config diff.ignoreSubmodules dirty
	git config submodule.fetchJobs 5
	git submodule update --init --jobs 5

# Update all submodules to latest remote head
.PHONY: update_submodules
update_submodules:
	git submodule update --remote --jobs 5

# Update all manually-enumerated submodules to latest remote head
.PHONY: update_submodules_force
update_submodules_force:
	git submodule foreach 'git -C $$toplevel submodule update --remote $$name'

.PHONY: list_submodules
list_submodules:
	@git ls-files --stage | grep ^160000 | cut -b51-
# Alternatively:
# git submodule foreach --quiet 'echo $name'

.PHONY: list_project_json
list_project_json:
	$(foreach project,$(PROJECTS),$(info $(project)))
	@:

.PHONY: list_dockerfiles
list_dockerfiles:
	$(foreach dockerfile,$(DOCKERFILES),$(info $(dockerfile)))
	@:

.PHONY: format_tools
format_tools:
	@underscore print -i tools/project.schema.json | tools/sponge tools/project.schema.json
	@underscore print -i .vscode/settings.json | tools/sponge .vscode/settings.json
	@underscore print -i .vscode/snippets.code-snippets | tools/sponge .vscode/snippets.code-snippets

.PHONY: todo
todo:
	@echo 'List programs:'
	@jq -r 'select(.programs==null and (.tags|contains(["programs"]))) | "- \(.id).json"' $(PROJECTS)
	@echo
	@echo 'Document Whitespace extension:'
	@jq -r 'select(.programs!=null and .whitespace.extension==null) | "- \(.id).json"' $(PROJECTS)
	@echo
	@echo 'Refine dates:'
	@jq -r 'select(.date | test("^\\d{4}$$"; "")) | "- \(.id).json: \(.date)"' $(PROJECTS)
	@echo
	@echo 'TODO:'
	@jq -r 'select(.todo != null) | "- \(.id).json: \(.todo)"' $(PROJECTS)
