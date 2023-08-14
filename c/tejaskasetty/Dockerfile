FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev flex-dev bison
WORKDIR /ws-compiler
COPY ws-compiler .
WORKDIR /ws-compiler/flex-bison
# Error: multiple definition of `symtable`
RUN flex lex.l && \
	  yacc --debug -d yacc.y && \
	  gcc -O3 -g -std=gnu11 -o ws -lfl lex.yy.c y.tab.c
WORKDIR /ws-compiler/flex-bison/262_267_256
RUN flex lex.l && \
	  yacc --debug -d yacc.y && \
	  gcc -O3 -g -std=gnu11 -o ws -lfl lex.yy.c y.tab.c

FROM alpine

RUN apk add --no-cache flex-libs
COPY --from=builder /ws-compiler/flex-bison/262_267_256/ws /
ENTRYPOINT ["/ws"]
