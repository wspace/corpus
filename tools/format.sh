#!/bin/sh

# Sort git submodules
# https://unix.stackexchange.com/questions/524048/sort-file-by-group-of-lines/524051#524051
perl -0ne 'print sort split(/^(?=\S)/m)' .gitmodules | sponge .gitmodules

# Format JSON and sort by date
# npm install -g underscore-cli
# https://github.com/ddopson/underscore-cli
underscore print -i projects.json |
  jq -c 'sort_by(.date) | reverse' |
  underscore print -o projects.json
echo >> projects.json
