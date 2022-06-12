FROM alpine as builder

RUN apk add git make gcc musl-dev
RUN git clone https://github.com/shinh/elvm
WORKDIR /elvm
RUN make
