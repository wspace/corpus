FROM alpine

RUN apk add git make gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/shinh/elvm
WORKDIR /home/elvm
RUN make
