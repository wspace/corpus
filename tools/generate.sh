#!/bin/sh

# Generate assembly.md

echo '# Whitespace assembly instructions

These are the names used by known Whitespace assembly dialects for
instructions, ranked by popularity.
' > assembly.md

jq -r '
reduce (.[].assembly.instructions | select (. != null) | to_entries[])
as $inst ({}; .[$inst.key] += $inst.value)
| to_entries[].value | map(ascii_downcase) | group_by(.) | sort_by(-length)
| map("`" + .[0] + "`" + (if length != 1 then " (" + (length | tostring) + ")" else "" end)) | "- " + join(", ")
' projects.json >> assembly.md
