FROM alpine

RUN apk add git gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/Sandeep023/Whitespace
WORKDIR /home/Whitespace
RUN gcc -O3 -Wall -o white y.tab.c lex.yy.c
RUN test -f /home/Whitespace/white
