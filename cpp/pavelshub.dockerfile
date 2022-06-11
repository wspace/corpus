FROM alpine

RUN apk add git make g++
WORKDIR /home
RUN git clone https://github.com/wspace/pavelshub-cpp wspace
WORKDIR /home/wspace
RUN make
RUN test -f /home/wspace/wspace
RUN test -f /home/wspace/assembler
