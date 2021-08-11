#!/bin/sh

# This lists the command lines used to build and run each project with
# the given file.

file="${1:-\$FILE}"

jq --arg file "$file" -r '
sort_by(.id)
| .[]
| (.run | .["interpret", "compile", "assemble"]) as $cmd
| select($cmd != null)
| [
    "cd \(.id)",
    .run.build // empty,
    ($cmd.usage
      | sub("\\$0"; "./" + $cmd.bin)
      | sub("<file>"; $file)
      | sub(" ?\\[[^\\]]+\\]"; ""; "g"))]
  | join(" && ")
' projects.json
