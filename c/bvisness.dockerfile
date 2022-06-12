FROM alpine as builder

RUN apk add git make gcc musl-dev flex bison
RUN git clone https://github.com/bvisness/whitespace
WORKDIR /whitespace
RUN make

FROM scratch as runner

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
