.PHONY: all
all:
	tools/format.sh
	go run tools/generate.go
