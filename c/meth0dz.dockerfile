FROM alpine

RUN apk add git gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/wspace/meth0dz-c whitespace
WORKDIR /home/whitespace
RUN gcc -O3 -Wall -o whitespace whitespace.c
RUN test -f /home/whitespace/whitespace
