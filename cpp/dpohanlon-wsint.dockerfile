FROM alpine

RUN apk add git make clang g++
WORKDIR /home
RUN git clone https://github.com/wspace/dpohanlon-wsint wsInt
WORKDIR /home/wsInt
RUN make
RUN clang++ -g -O2 -Wall -pedantic -o bin/ParserTest tests/ParserTest.cpp src/parser/*.cpp
RUN bin/ParserTest
RUN test -f /home/wsInt/bin/wsInt
