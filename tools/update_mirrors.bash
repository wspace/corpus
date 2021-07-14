#!/bin/bash -e

# Updates mirrors with changes from upstream

mkdir -p tmp && cd tmp

git clone https://github.com/hostilefork/whitespacers akers-lolcode
git -C akers-lolcode filter-repo --subdirectory-filter lolcode
git -C akers-lolcode branch -m master main
git -C akers-lolcode remote add origin https://github.com/wspace/akers-lolcode

git clone https://github.com/Aniket965/Hello-world eah-hello-world
# Ignore dirty tree from conflicting case-sensitive files with --force
git -C eah-hello-world filter-repo --force \
  --subdirectory-filter Whitespace \
  --path LICENSE
git -C eah-hello-world branch -m master main
git -C eah-hello-world remote add origin https://github.com/wspace/eah-hello-world

git clone https://github.com/hogelog/hogel.org-old hogelog-cpp
git -C hogelog-cpp filter-repo \
  --subdirectory-filter content/lib/c/ws \
  --path content/lib/c/ws_20080502.tar.bz2 \
  --path content/lib/c/ws_20080503.tar.bz2 \
  --path-rename content/lib/c/:
git -C hogelog-cpp branch -m master main
git -C hogelog-cpp remote add origin https://github.com/wspace/hogelog-cpp

# Pushes cleanly, if no history was rewritten
git -C akers-lolcode push -u origin main
git -C eah-hello-world push -u origin main
git -C hogelog-cpp push -u origin main
