#!/bin/sh -e

src="${1:?usage <src> <dest>}"
dest="${2:?usage <src> <dest>}"

git mv "$src" "$dest"
git mv "$src.json" "$dest.json"
jq -c --arg dest "$dest" '.id = $dest' "$dest.json" | underscore print | sponge "$dest.json"
mv ".git/modules/$src" ".git/modules/$dest"
echo "gitdir: ../../.git/modules/$dest" > "$dest/.git"
sed -i '' "s,^\[submodule \"$src\"\]$,[submodule \"$dest\"]," .gitmodules
tools/format_gitmodules.sh
git add .gitmodules
git submodule init "$dest"
