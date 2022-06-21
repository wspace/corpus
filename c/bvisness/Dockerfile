FROM alpine AS builder

RUN apk add make gcc musl-dev flex bison
WORKDIR /whitespace
COPY whitespace .
RUN make

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
