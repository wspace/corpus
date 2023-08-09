FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev
WORKDIR /Whitespace
COPY Whitespace .
RUN make DIFF='diff -b' test clean
RUN gcc -O3 -g -static \
        -DNDEBUG \
        -DMAX_SOURCE_SIZE=65536 \
        -DMAX_BYTECODE_SIZE=1048576 \
        -DMAX_LABEL_LENGTH=65536 \
        -DMAX_N_LABEL=1024 \
        -DUNDEF_LIST_SIZE=256 \
        -DSTACK_SIZE=65536 \
        -DHEAP_SIZE=65536 \
        -DCALL_STACK_SIZE=65536 \
        -DWS_INT=int \
        -DWS_ADDR_INT='unsigned int' \
        -DINDENT_STR='"  "' \
        -o whitespace.out whitespace.c

FROM scratch

COPY --from=builder /Whitespace/whitespace.out /
COPY Whitespace/t/*.ws /
ENTRYPOINT ["/whitespace.out"]
