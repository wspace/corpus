#!/bin/bash -e

files=()
for file in "${@:1}"; do
  files+=("$file" sep)
done
cat "${files[@]}" | gows build/whitespace.ws
