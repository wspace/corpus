FROM alpine AS builder

RUN apk add make clang g++
WORKDIR /wsInt
COPY wsInt .
RUN make CXX='clang++ -static'
RUN clang++ -g -O2 -o bin/ParserTest tests/ParserTest.cpp src/parser/*.cpp
RUN bin/ParserTest

FROM scratch

COPY --from=builder /wsInt/bin/wsInt /
ENTRYPOINT ["/wsInt"]
