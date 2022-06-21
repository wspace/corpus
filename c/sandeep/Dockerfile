FROM alpine AS builder

RUN apk add git gcc musl-dev
RUN git clone https://github.com/Sandeep023/Whitespace
WORKDIR /Whitespace
RUN gcc -O3 -Wall -static -o white y.tab.c lex.yy.c

FROM scratch

COPY --from=builder /Whitespace/white /
ENTRYPOINT ["/white"]
