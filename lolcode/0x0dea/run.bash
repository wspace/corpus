#!/bin/bash -e

if [ $# != 1 ]; then
  echo 'Usage: whitespace.lol <file>' >&2
  exit 2
fi
file="$1"

# Echo filename, then pass stdin:
cat <(echo "$file") - | lci whitespace.lol

# If the Whitespace program contains no read instructions, it can be run
# without stdin:
# lci whitespace.lol <<< "$file"
