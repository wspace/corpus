FROM alpine

RUN apk add git g++
WORKDIR /home
RUN git clone https://github.com/noia1/Whitespace-Interpreter
WORKDIR /home/Whitespace-Interpreter
RUN g++ -O3 -Wall -Werror -o ws Lexer.cc Parser.cc
RUN test -f /home/Whitespace-Interpreter/ws
