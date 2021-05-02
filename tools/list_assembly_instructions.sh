#!/bin/sh

jq 'reduce (.[].assembly.instructions | select (. != null) | to_entries[]) as $inst ({}; .[$inst.key] = (.[$inst.key] + $inst.value | sort))' projects.json
