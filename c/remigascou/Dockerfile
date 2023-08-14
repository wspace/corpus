FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev flex-dev bison
WORKDIR /whitespace
COPY whitespace .
WORKDIR /whitespace/dev
RUN mkdir bin && \
    \
    flex -o src/compiler/lex.yy.c src/compiler/lexicograph.l && \
    yacc -d -o src/compiler/y.tab.c src/compiler/compiler.y && \
    gcc -O3 -g -o bin/wscc -I src/lib/ -lfl -ly src/wscc.c src/compiler/lex.yy.c src/compiler/y.tab.c src/lib/labelparser.c && \
    \
    flex -o src/decompiler/lex.yy.c src/decompiler/lexicograph.l && \
    yacc -d -o src/decompiler/y.tab.c src/decompiler/decompiler.y && \
    gcc -O3 -g -o bin/decompiler -I src/lib/ -lfl -ly src/decompiler/lex.yy.c src/decompiler/y.tab.c src/lib/labelparser.c && \
    \
    sh tests.sh

# lexxyacc/compiler/compiler.y requires LList and LList_asm types from missing
# local headers, so cannot be built.

FROM alpine

RUN apk add --no-cache flex-libs
COPY --from=builder /whitespace/dev/bin/wscc /whitespace/dev/bin/decompiler /
COPY whitespace/testprograms /testprograms
ENTRYPOINT ["/wscc"]
