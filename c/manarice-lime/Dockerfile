FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev patch
WORKDIR /utils/limelib
COPY limelib 0001-Fix-double-free-and-missing-frees.patch ./
RUN patch -p1 -i 0001-Fix-double-free-and-missing-frees.patch && \
    gcc -g -o test test.c src/*.c && \
    ./test
WORKDIR /utils/whitespace
COPY whitespace .
RUN gcc -O3 -g -static -o lwsvm whitespace_vm.c ../limelib/src/*.c && \
    gcc -O3 -g -static -o lwsa wsa.c ../limelib/src/*.c

FROM scratch

COPY --from=builder /utils/whitespace/lwsvm /utils/whitespace/lwsa /
COPY whitespace/ws /ws/
ENTRYPOINT ["/lwsvm"]
