FROM alpine

RUN apk add git make gcc musl-dev flex
WORKDIR /home
RUN git clone https://github.com/wspace/rdebath-c whitespace
WORKDIR /home/whitespace
RUN make
RUN test -f /home/whitespace/ws2c
RUN test -f /home/whitespace/wsa
