FROM alpine

RUN apk add git make gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/wspace/koturn-c Whitespace
WORKDIR /home/Whitespace
RUN make
RUN test -f /home/Whitespace/whitespace.out
