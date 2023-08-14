FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /whitespace
COPY whitespace .
RUN mkdir bin && \
    gcc -O3 -g -static -o bin/wsvm src/wsvm.c src/num2space.c src/lexer.c src/fun_stack.c src/fun_heap_access.c src/fun_flow_ctrl.c src/fun_io.c && \
    gcc -O3 -g -static -o bin/wscc src/wscc.c src/debug.c src/lexer.c src/num2space.c && \
    gcc -O3 -g -static -o bin/wsas src/wsas.c && \
    gcc -O3 -g -static -o bin/ws_trans src/ws_trans.c && \
    gcc -O3 -g -static -o bin/wb_dump src/wb_dump.c src/debug.c

FROM scratch

COPY --from=builder /whitespace/bin /
ENTRYPOINT ["/wsvm"]
