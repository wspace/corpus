#!/bin/sh -e

src="${1:?usage <src> <dest>}"
dest="${2:?usage <src> <dest>}"
src="${src%%/}"
dest="${dest%%/}"

git mv "$src" "$dest"
jq -c --arg dest "$dest" '.id = $dest' "$dest/project.json" | underscore print | sponge "$dest/project.json"
mv ".git/modules/$src" ".git/modules/$dest"
echo "gitdir: ../../../.git/modules/$dest" > "$dest/.git"
sed -i '' "s,^\[submodule \"$src/,[submodule \"$dest/," .gitmodules
tools/format_gitmodules.sh
git add .gitmodules
find "$dest" -type d -depth 1 -exec git submodule init {} \;
