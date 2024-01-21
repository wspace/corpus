# Generate missing_submodules.md

"# Missing submodules

<!-- Generated by tools/generate_missing_submodules.jq; DO NOT EDIT. -->

These projects are missing a submodule, usually because I have not yet converted
the source to git.
",
(
  .[] | select((.submodules | length > 0) or .source_unavailable | not) |
  "- \(.id)"
)