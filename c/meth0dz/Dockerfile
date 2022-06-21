FROM alpine AS builder

RUN apk add gcc musl-dev
WORKDIR /whitespace
COPY whitespace .
RUN gcc -O3 -Wall -static -o whitespace whitespace.c

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
