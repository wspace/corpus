#!/bin/sh

# Sort git submodules
# https://unix.stackexchange.com/questions/524048/sort-file-by-group-of-lines/524051#524051

perl -0ne 'print sort split(/^(?=\S)/m)' .gitmodules > tmp
mv tmp .gitmodules
