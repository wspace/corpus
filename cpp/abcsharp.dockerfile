FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/abcsharp/Whitespace
WORKDIR /home/Whitespace
RUN g++ -O3 -Wall -std=c++11 -o wsi whitespace/main.cpp
RUN test -f /home/Whitespace/wsi
