FROM alpine AS builder

RUN apk add make gcc musl-dev
WORKDIR /elvm
COPY elvm .
RUN make
