#!/bin/sh -e

dockerfiles=${*:-*/Dockerfile */*/Dockerfile}

echo 'docker build -q -t wspace-corpus/haskell/edwinb-wspace-0.3 --platform linux/amd64 haskell/edwinb-wspace-0.3'
echo 'docker run -it --rm wspace-corpus/haskell/edwinb-wspace-0.3 hworld.ws'
echo

get_id() {
  id="${1%/Dockerfile}"
  id="${id%/project.json}"
  echo "${id%/}"
}

for f in $dockerfiles; do
  id="$(get_id "$f")"
  echo "docker build -q -t wspace-corpus/$id $id"
done
echo
for f in $dockerfiles; do
  id="$(get_id "$f")"
  echo "docker run -it --rm wspace-corpus/$id"
done
