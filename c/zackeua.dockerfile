FROM alpine

RUN apk add git gcc g++
WORKDIR /home
RUN git clone https://github.com/zackeua/Whitespace
WORKDIR /home/Whitespace
RUN gcc -g -Wall -o wlang whitespace.c
RUN g++ -g -Wall -o whitespacecpp whitespace.cpp
RUN test -f /home/Whitespace/wlang
RUN test -f /home/Whitespace/whitespacecpp
