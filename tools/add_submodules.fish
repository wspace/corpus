#!/usr/bin/env fish

# Add submodules for all newly-added projects

set projects (jq -j '
map(select((.path|length) != 0 and (.source[0] | test("^https://github.com/[^/]+/[^/]+$")))) |
sort_by(.path)[] | "\(.path)|\(.source[0])\n"
' projects.json)

for p in $projects
  set p (string split '|' $p)
  if test ! -e $p[1]
    git submodule add $p[2] $p[1]
  end
end
