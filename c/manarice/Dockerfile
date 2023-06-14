FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev
WORKDIR /utils/limelib
COPY limelib .
RUN gcc -O3 -static -o test test.c src/*.c
RUN ./test
WORKDIR /utils/whitespace
COPY whitespace .
RUN gcc -O3 -static -o lwsvm whitespace_vm.c ../limelib/src/*.c && \
    gcc -O3 -static -o lwsa wsa.c ../limelib/src/*.c

FROM scratch

COPY --from=builder /utils/whitespace/lwsvm /
COPY --from=builder /utils/whitespace/lwsa /
COPY whitespace/ws /ws/
ENTRYPOINT ["/lwsvm"]
