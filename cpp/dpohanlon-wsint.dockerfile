FROM alpine as builder

RUN apk add git make clang g++
RUN git clone https://github.com/wspace/dpohanlon-wsint wsInt
WORKDIR /wsInt
RUN make
RUN clang++ -g -O2 -Wall -pedantic -o bin/ParserTest tests/ParserTest.cpp src/parser/*.cpp
RUN bin/ParserTest

FROM scratch as runner

COPY --from=builder /wsInt/bin/wsInt /
ENTRYPOINT ["/wsInt"]
