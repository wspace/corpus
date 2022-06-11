FROM alpine

RUN apk add git make g++
WORKDIR /home
RUN git clone https://github.com/andrewarchi/respace
WORKDIR /home/respace
RUN make test
RUN make
RUN test -f /home/respace/respace
