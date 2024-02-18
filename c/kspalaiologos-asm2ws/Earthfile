FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev flex bison
WORKDIR /asm2ws
COPY asm2ws .
RUN ./configure --with-target=release && \
    make -j$(nproc) wsi CFLAGS='-O3 -g -static'

FROM scratch

COPY --from=builder /asm2ws/wsi /
COPY asm2ws/tests /tests
ENTRYPOINT ["/wsi"]
