FROM alpine as builder

RUN apk add git gcc musl-dev
RUN git clone https://github.com/nobody1986/whitespace-php
WORKDIR /whitespace-php
RUN gcc -O3 -o whitespace whitespace.c list.c stack.c

FROM scratch as runner

COPY --from=builder /whitespace-php/whitespace /
ENTRYPOINT ["/whitespace"]
