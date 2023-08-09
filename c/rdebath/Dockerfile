FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev flex
WORKDIR /whitespace
COPY whitespace .
RUN cp wsc.sh wsc && \
    \
    gcc -O3 -g -static -o blockquote blockquote.c && \
    ( echo 'char * header =' && \
      ./blockquote -f ws_engine.h && \
      echo ';' ) > ws_engine_txt.h && \
    gcc -O3 -g -static -o ws2c ws2c.c && \
    \
    flex -t wsa.l > wsa.c && \
    gcc -O3 -g -static -o wsa wsa.c

FROM scratch

COPY --from=builder /whitespace/wsc /whitespace/ws2c /whitespace/wsa /
COPY whitespace/tests /tests
ENTRYPOINT ["/ws2c"]
