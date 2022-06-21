FROM alpine AS builder

RUN apk add gcc g++
WORKDIR /Whitespace
COPY Whitespace .
RUN gcc -g -Wall -static -o wlang whitespace.c
RUN g++ -g -Wall -static -o whitespacecpp whitespace.cpp

FROM scratch

COPY --from=builder /Whitespace/wlang /
COPY --from=builder /Whitespace/whitespacecpp /
ENTRYPOINT ["/wlang"]
