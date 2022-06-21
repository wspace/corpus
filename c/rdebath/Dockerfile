FROM alpine AS builder

RUN apk add make gcc musl-dev flex
WORKDIR /whitespace
COPY whitespace .
RUN make

FROM scratch

COPY --from=builder /whitespace/ws2c /
COPY --from=builder /whitespace/wsa /
ENTRYPOINT ["/ws2c"]
