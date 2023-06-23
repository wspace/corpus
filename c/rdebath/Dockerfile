FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev flex
WORKDIR /whitespace
COPY whitespace .
RUN make

FROM scratch

COPY --from=builder /whitespace/ws2c /
COPY --from=builder /whitespace/wsa /
COPY whitespace/tests /tests
ENTRYPOINT ["/ws2c"]
