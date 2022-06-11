FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/Keirua/whitespace
WORKDIR /home/whitespace
RUN g++ -O3 -Wall -o white main.cpp
RUN test -f /home/whitespace/white
