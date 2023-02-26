FROM alpine AS builder

RUN apk add --no-cache gcc g++
WORKDIR /Whitespace
COPY Whitespace .
RUN gcc -O3 -static -o wlang whitespace.c
RUN g++ -O3 -static -o whitespacecpp whitespace.cpp

FROM scratch

COPY --from=builder /Whitespace/wlang /
COPY --from=builder /Whitespace/whitespacecpp /
ENTRYPOINT ["/wlang"]
