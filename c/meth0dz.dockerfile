FROM alpine as builder

RUN apk add git gcc musl-dev
RUN git clone https://github.com/wspace/meth0dz-c whitespace
WORKDIR /whitespace
RUN gcc -O3 -Wall -static -o whitespace whitespace.c

FROM scratch as runner

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
