FROM alpine

RUN apk add git make g++
WORKDIR /home
RUN git clone https://github.com/drebelsky/whitespace-jit
WORKDIR /home/whitespace-jit
RUN make CXXFLAGS='-O3 -Wall -Wpedantic -std=c++17'
RUN test -f /home/whitespace-jit/compile
