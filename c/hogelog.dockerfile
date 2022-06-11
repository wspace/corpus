FROM alpine

RUN apk add git make gcc musl-dev perl
WORKDIR /home
RUN git clone https://github.com/wspace/hogelog-c ws
WORKDIR /home/ws
RUN perl parsegen.pl parse.def > gencode.c
RUN make
RUN test -f /home/ws/wspace
