FROM alpine as builder

RUN apk add git gcc g++
RUN git clone https://github.com/zackeua/Whitespace
WORKDIR /Whitespace
RUN gcc -g -Wall -static -o wlang whitespace.c
RUN g++ -g -Wall -static -o whitespacecpp whitespace.cpp

FROM scratch as runner

COPY --from=builder /Whitespace/wlang /
COPY --from=builder /Whitespace/whitespacecpp /
ENTRYPOINT ["/wlang"]
