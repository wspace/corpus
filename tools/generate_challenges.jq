#!/bin/bash
# Usage: generate_challenges.jq [project.json]... > challenges.md \
exec jq -rsf "$0" -- "${@:1}"

# Generate challenges.md

"# Whitespace Coding challenges

<!-- Generated by tools/generate_challenges.jq; DO NOT EDIT. -->

These are projects that contain solutions for coding challenges related to
Whitespace. These can be solutions written in Whitespace or, in the case of
Codewars, interpreters for Whitespace.",
(
  reduce (.[] | select(.challenges != null)) as $project ({};
    reduce $project.challenges[] as $challenge (.;
      .[$challenge] += [$project.id])) |
  to_entries | sort_by(.key | ascii_downcase)[] |
  (
    "",
    "## \(.key)",
    "",
    (.value | sort[] | "- [\(.)](\(.))")
  )
)
