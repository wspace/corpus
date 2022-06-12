FROM alpine as builder

RUN apk add git make gcc musl-dev flex bison
RUN git clone https://github.com/kspalaiologos/asm2ws
WORKDIR /asm2ws
RUN ./configure --with-target=release
RUN make -j4 wsi

FROM scratch as runner

COPY --from=builder /asm2ws/wsi /
ENTRYPOINT ["/wsi"]
