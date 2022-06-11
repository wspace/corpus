FROM alpine

RUN apk add git make g++ gmp-dev
WORKDIR /home
RUN git clone https://git.code.sf.net/p/esco/code esco
WORKDIR /home/esco
RUN ./configure && make
RUN test -f /home/esco/src/esco
