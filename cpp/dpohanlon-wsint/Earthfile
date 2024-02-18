FROM alpine AS builder

RUN apk add --no-cache make clang g++
WORKDIR /wsInt
COPY wsInt .
RUN make CPPFLAGS='-O3 -static'
RUN clang++ -O3 -o bin/ParserTest tests/ParserTest.cpp src/parser/*.o
RUN bin/ParserTest

FROM scratch

COPY --from=builder /wsInt/bin/wsInt /
ENTRYPOINT ["/wsInt"]
