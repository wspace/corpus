FROM alpine AS builder

RUN apk add gcc musl-dev
WORKDIR /Whitespace
COPY Whitespace .
RUN gcc -O3 -Wall -static -o white y.tab.c lex.yy.c

FROM scratch

COPY --from=builder /Whitespace/white /
ENTRYPOINT ["/white"]
