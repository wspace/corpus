FROM alpine AS builder

RUN apk add gcc musl-dev
WORKDIR /whitespace-php
COPY whitespace-php .
RUN gcc -O3 -o whitespace whitespace.c list.c stack.c

FROM scratch

COPY --from=builder /whitespace-php/whitespace /
ENTRYPOINT ["/whitespace"]
