# List `docker build` commands

"docker build -t wspace-corpus/crates-io -f tools/docker/crates-io.dockerfile /dev/null",
(.[] | select(.build != null) | "docker build -t wspace-corpus/\(.id) -f \(.id).dockerfile /dev/null")
