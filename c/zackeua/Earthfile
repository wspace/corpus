FROM alpine AS builder

WORKDIR /Whitespace
COPY Whitespace .
RUN <<EOF
set -e

apk add --no-cache gcc g++
gcc -O3 -static -o wlang whitespace.c
g++ -O3 -static -o whitespacecpp whitespace.cpp

EOF

FROM scratch

COPY --from=builder /Whitespace/wlang /
COPY --from=builder /Whitespace/whitespacecpp /
ENTRYPOINT ["/wlang"]
