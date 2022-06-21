FROM alpine AS builder

RUN apk add git gcc musl-dev
RUN git clone https://github.com/wspace/meth0dz-c whitespace
WORKDIR /whitespace
RUN gcc -O3 -Wall -static -o whitespace whitespace.c

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
