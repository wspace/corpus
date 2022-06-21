FROM alpine AS builder

RUN apk add git make gcc musl-dev flex bison
RUN git clone https://github.com/wspace/remigascou-c whitespace
WORKDIR /whitespace/dev
RUN make
WORKDIR /whitespace/lexxyacc
RUN make

FROM scratch

COPY --from=builder /whitespace/dev/bin/compiler /
COPY --from=builder /whitespace/dev/bin/decompiler /
COPY --from=builder /whitespace/lexxyacc/compiler/compiler /
ENTRYPOINT ["/compiler"]
