FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/kapustaikwas27/Whitespace-compiler
WORKDIR /home/Whitespace-compiler
RUN g++ -O3 -Wall -o pre pre.cpp
RUN test -f /home/Whitespace-compiler/pre
