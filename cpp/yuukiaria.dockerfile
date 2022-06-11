FROM alpine

RUN apk add git make g++
WORKDIR /home
RUN git clone https://github.com/YuukiARIA/ws-interpreter
WORKDIR /home/ws-interpreter/src
RUN mkdir ../bin
RUN make TARGET=../bin/ws
RUN test -f /home/ws-interpreter/bin/ws
