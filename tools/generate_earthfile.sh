#!/usr/bin/env bash
set -eEuo pipefail

lang="$1"
projects=("${@:2}")

echo 'VERSION 0.8'
echo
for project in "${projects[@]}"; do
  echo "IMPORT ./$project AS $lang-$project"
done
echo
echo 'build:'
echo '    FROM scratch'
for project in "${projects[@]}"; do
  echo "    BUILD $lang-$project+build"
done
echo "    WORKDIR /corpus/$lang"
for project in "${projects[@]}"; do
  echo "    COPY $lang-$project+build/ $project/"
done
echo '    SAVE ARTIFACT /corpus /corpus'
echo
echo 'docker:'
echo '    FROM alpine'
echo "    WORKDIR /corpus/$lang"
echo '    COPY +build/ /'
echo "    SAVE IMAGE wspace-corpus/$lang"
echo
echo 'docker-all:'
echo '    BUILD +docker'
for project in "${projects[@]}"; do
  echo "    BUILD $lang-$project+docker"
done
