FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /Whitespace
COPY Whitespace .
RUN gcc -O3 -g -static -o white y.tab.c lex.yy.c

FROM scratch

COPY --from=builder /Whitespace/white /
COPY Whitespace/*.ws /
ENTRYPOINT ["/white"]
