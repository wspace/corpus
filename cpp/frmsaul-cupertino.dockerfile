FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/frmsaul/Cupertino-WhiteSpace-Interperter
WORKDIR /home/Cupertino-WhiteSpace-Interperter
RUN g++ -O3 -o whitespace src/*.cpp
RUN test -f /home/Cupertino-WhiteSpace-Interperter/whitespace
