FROM alpine AS builder

RUN apk add make gcc musl-dev flex bison
WORKDIR /asm2ws
COPY asm2ws .
RUN ./configure --with-target=release
RUN make -j4 wsi

FROM scratch

COPY --from=builder /asm2ws/wsi /
ENTRYPOINT ["/wsi"]
