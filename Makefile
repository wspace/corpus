.PHONY: generate
generate:
	go run tools/generate.go
	tools/sort_gitmodules.sh
