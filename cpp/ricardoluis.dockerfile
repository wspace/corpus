FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/RicardoLuis0/whitespace
WORKDIR /home/whitespace
RUN g++ -O3 -Wall -std=c++11 -o whitespace whitespace.cpp
RUN test -f /home/whitespace/whitespace
