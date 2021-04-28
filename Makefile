.PHONY: all
all:
	go run tools/generate.go
	tools/format.sh
