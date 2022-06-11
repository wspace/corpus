FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/t3nsor/SPOJ
WORKDIR /home/SPOJ
RUN g++ -g -std=c++17 -o $0 wstp.cpp
RUN test -f /home/SPOJ/wstp
