FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/sudheesh4/Whitespace-Interpreter-C-
WORKDIR /home/Whitespace-Interpreter-C-
RUN g++ -O3 -o space space.cpp
RUN test -f /home/Whitespace-Interpreter-C-/space
