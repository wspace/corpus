#!/bin/sh -e

# List `docker build` commands
for f in crates-io lci; do
  echo "docker build -t wspace-corpus/$f -f tools/docker/$f.dockerfile /dev/null"
done
for f in */*.dockerfile; do
  f="${f%.dockerfile}"
  echo "docker build -t wspace-corpus/$f -f $f.dockerfile /dev/null"
done
