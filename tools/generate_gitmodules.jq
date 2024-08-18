#!/bin/bash
# Usage: generate_gitmodules.jq [project.json]... > .gitmodules \
exec jq -rsf "$0" -- "${@:1}"

# Generate .gitmodules

def escape_space:
  if . | contains(" ") then "\(.)" else . end;

[
  .[] |
  .id as $id |
  .submodules[]? |
  .path = "\($id)/\(.path)"
] + [
  {
	  path: "collections/hostilefork-whitespacers",
	  url: "https://github.com/hostilefork/whitespacers",
  },
  {
	  path: "collections/thaliaarchi-ws-archive",
	  url: "https://github.com/wspace/ws-archive",
  }
] |
sort_by(.path)[] |

"[submodule \"\(.path)\"]
\tpath = \(.path | escape_space)
\turl = \(.url)",
if .branch != null then "\tbranch = \(.branch)" else empty end
