# Add submodules for all newly-added projects

map(select((.path|length) != 0 and (.source[0] | test("^https://github.com/[^/]+/[^/]+$")))) |
sort_by(.path)[] |
"[ -e '\(.path)' ] || git submodule add '\(.source[0])' '\(.path)'"
