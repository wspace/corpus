FROM alpine

RUN apk add git make gcc musl-dev flex bison
WORKDIR /home
RUN git clone https://github.com/bvisness/whitespace
WORKDIR /home/whitespace
RUN make
RUN test -f /home/whitespace/whitespace
