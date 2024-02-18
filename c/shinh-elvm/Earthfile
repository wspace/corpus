FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev
WORKDIR /elvm
COPY elvm .
RUN make
