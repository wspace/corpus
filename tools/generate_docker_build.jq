# List `docker build` commands

("crates-io", "lci" | "docker build -t wspace-corpus/\(.) -f tools/docker/\(.).dockerfile /dev/null"),
(.[] | select(.build != null) | "docker build -t wspace-corpus/\(.id) -f \(.id).dockerfile /dev/null")
