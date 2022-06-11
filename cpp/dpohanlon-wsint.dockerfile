FROM alpine

RUN apk add git make clang g++
WORKDIR /home
RUN git clone https://github.com/wspace/dpohanlon-wsint wsInt
WORKDIR /home/wsInt
RUN make
RUN test -f /home/wsInt/bin/wsInt
