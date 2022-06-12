#!/bin/sh -e

echo 'docker build -t wspace-corpus/haskell/edwinb-wspace-0.3 -f haskell/edwinb-wspace-0.3.dockerfile --platform linux/amd64 /dev/null'
echo 'docker run -i -t --rm wspace-corpus/haskell/edwinb-wspace-0.3 hworld.ws'

# List `docker build` commands
for f in crates-io lci; do
  echo "docker build -t wspace-corpus/$f -f tools/docker/$f.dockerfile /dev/null"
  echo "docker run -i -t --rm wspace-corpus/$f"
done
for f in */*.dockerfile; do
  f="${f%.dockerfile}"
  echo "docker build -t wspace-corpus/$f -f $f.dockerfile /dev/null"
  echo "docker run -i -t --rm wspace-corpus/$f"
done
