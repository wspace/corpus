FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/wspace/peasley-cpp whitespace
WORKDIR /home/whitespace
RUN g++ -O3 -Wall -o whitespace whitespace.cpp
RUN test -f /home/whitespace/whitespace
