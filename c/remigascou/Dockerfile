FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev flex bison
WORKDIR /whitespace
COPY whitespace .
RUN make -C dev && \
    make -C lexxyacc

FROM scratch

COPY --from=builder /whitespace/dev/bin/compiler \
                    /whitespace/dev/bin/decompiler \
                    /whitespace/lexxyacc/compiler/compiler /
COPY whitespace/testprograms /testprograms
ENTRYPOINT ["/compiler"]
