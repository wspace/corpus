FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/abac00s/WHINT
WORKDIR /home/WHINT
RUN g++ -O3 -Wall -o whint whint.cpp
RUN test -f /home/WHINT/whint
