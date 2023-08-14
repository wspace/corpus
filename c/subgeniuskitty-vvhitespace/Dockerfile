FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev python3
WORKDIR /vvhitespace
COPY vvhitespace .
RUN gcc -O3 -g -static -std=c99 -o vvc vv_compiler.c && \
    gcc -O3 -g -static -std=c99 -o vvi vv_interpreter.c && \
    (cd tests && python3 vv_test.py) && \
    (cd stdlib_tests && python3 vv_test.py)

FROM scratch

COPY --from=builder /vvhitespace/vvc /vvhitespace/vvi /
ENTRYPOINT ["/vvi"]
