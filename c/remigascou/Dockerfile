FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev flex bison
WORKDIR /whitespace
COPY whitespace .
RUN make -C dev
RUN make -C lexxyacc

FROM scratch

COPY --from=builder /whitespace/dev/bin/compiler /
COPY --from=builder /whitespace/dev/bin/decompiler /
COPY --from=builder /whitespace/lexxyacc/compiler/compiler /
ENTRYPOINT ["/compiler"]
