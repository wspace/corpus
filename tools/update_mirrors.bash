#!/bin/bash -e

# Updates mirrors with changes from upstream

mkdir -p tmp && cd tmp

git clone https://github.com/hostilefork/whitespacers
cp -rp whitespacers hostilefork-rebol
cp -rp whitespacers hostilefork-ren-c
mv whitespacers akers-lolcode
git -C hostilefork-rebol filter-repo --subdirectory-filter rebol
git -C hostilefork-ren-c filter-repo --subdirectory-filter ren-c
git -C akers-lolcode filter-repo --subdirectory-filter lolcode

git clone https://github.com/Aniket965/Hello-world eah-hello-world
git -C eah-hello-world filter-repo --force --subdirectory-filter Whitespace --path LICENSE

git -C hostilefork-rebol branch -m master main
git -C hostilefork-ren-c branch -m master main
git -C akers-lolcode branch -m master main
git -C eah-hello-world branch -m master main

git -C hostilefork-rebol remote add origin https://github.com/wspace/hostilefork-rebol
git -C hostilefork-ren-c remote add origin https://github.com/wspace/hostilefork-ren-c
git -C akers-lolcode remote add origin https://github.com/wspace/akers-lolcode
git -C eah-hello-world remote add origin https://github.com/wspace/eah-hello-world

# Pushes cleanly, if no history was rewritten
git -C hostilefork-rebol push -u origin main
git -C hostilefork-ren-c push -u origin main
git -C akers-lolcode push -u origin main
git -C eah-hello-world push -u origin main
