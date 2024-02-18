FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev flex bison
WORKDIR /whitespace
COPY whitespace .
RUN flex -o whitespace.yy.c whitespace.l && \
    bison -v -d whitespace.y && \
    gcc -O3 -g -static -o whitespace whitespace.yy.c whitespace.tab.c parsetree.c stack.c heap.c linkedlist.c && \
    gcc -g -o stacktest tests/stacktest.c stack.c linkedlist.c && \
    gcc -g -o heaptest tests/heaptest.c heap.c linkedlist.c && \
    gcc -g -o linkedlisttest tests/linkedlisttest.c linkedlist.c
RUN ./stacktest && \
    ./heaptest && \
    ./linkedlisttest

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
