#!/bin/bash -e

# Updates mirrors with changes from upstream

mkdir -p tmp && cd tmp

git clone https://github.com/hostilefork/whitespacers akers-lolcode
git -C akers-lolcode filter-repo --subdirectory-filter lolcode
git -C akers-lolcode branch -m master main
git -C akers-lolcode remote add origin https://github.com/wspace/akers-lolcode

git clone https://github.com/Aniket965/Hello-world eah-hello-world
git -C eah-hello-world filter-repo --force --subdirectory-filter Whitespace --path LICENSE
git -C eah-hello-world branch -m master main
git -C eah-hello-world remote add origin https://github.com/wspace/eah-hello-world

# Pushes cleanly, if no history was rewritten
git -C akers-lolcode push -u origin main
git -C eah-hello-world push -u origin main
