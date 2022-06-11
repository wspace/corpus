FROM alpine

RUN apk add git gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/nobody1986/whitespace-php
WORKDIR /home/whitespace-php
RUN gcc -O3 -o whitespace whitespace.c list.c stack.c
RUN test -f /home/whitespace-php/whitespace
