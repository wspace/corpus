# Add submodules for all newly-added projects

map(select((.path|length) != 0 and (.source[0] | test("^https://((gist\\.)?github|gitlab)\\.com/[^/]+/[^/]+$")))) |
sort_by(.path)[] |
(.source[0] | if startswith("https://gitlab.com/") then . + ".git" else . end) as $src |
"[ -e '\(.path)' ] || git submodule add '\($src)' '\(.path)'"
