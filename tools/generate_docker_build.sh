#!/bin/sh -e

echo 'docker build -t wspace-corpus/haskell/edwinb-wspace-0.3 -f haskell/edwinb-wspace-0.3.dockerfile --platform linux/amd64 /dev/null'
echo 'docker run -i -t --rm wspace-corpus/haskell/edwinb-wspace-0.3 hworld.ws'
echo 'docker build -t wspace-corpus/rust -f tools/rust/Dockerfile /dev/null'
echo 'docker run -i -t --rm wspace-corpus/rust'

for f in */*.dockerfile; do
  f="${f%.dockerfile}"
  echo "docker build -t wspace-corpus/$f -f $f.dockerfile /dev/null"
  echo "docker run -i -t --rm wspace-corpus/$f"
done
