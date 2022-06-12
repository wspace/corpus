#!/bin/sh -e

dockerfiles=${*:-*/Dockerfile */*.dockerfile}

echo 'docker build -t wspace-corpus/haskell/edwinb-wspace-0.3 -f haskell/edwinb-wspace-0.3.dockerfile --platform linux/amd64 /var/empty'
echo 'docker run -i -t --rm wspace-corpus/haskell/edwinb-wspace-0.3 hworld.ws'

get_tag() {
  t="${1%.dockerfile}"
  t="${t%/Dockerfile}"
  echo "wspace-corpus/$t"
}

for f in $dockerfiles; do
  echo "docker build -q -t $(get_tag "$f") -f $f /var/empty"
done
for f in $dockerfiles; do
  echo "docker run -i -t --rm $(get_tag "$f")"
done
