FROM alpine AS builder

WORKDIR /whitespace
COPY whitespace .
RUN <<EOF
set -e

apk add --no-cache gcc musl-dev flex

cp wsc.sh wsc

gcc -O3 -g -static -o blockquote blockquote.c
( echo 'char * header ='
  ./blockquote -f ws_engine.h
  echo ';' ) > ws_engine_txt.h
gcc -O3 -g -static -o ws2c ws2c.c

flex -t wsa.l > wsa.c
gcc -O3 -g -static -o wsa wsa.c

gcc -o gencode-deadfish-in-ws gencode-deadfish-in-ws.c
gcc -o gen-99bottles gen-99bottles.c
gcc -o gen-prime gen-prime.c
gcc -o gen-wsinterws gen-wsinterws.c
./gencode-deadfish-in-ws > deadfish.ws
./gen-99bottles > 99bottles.ws
./gen-prime > prime.ws
./gen-wsinterws > wsinterws.ws

EOF

FROM scratch

COPY --from=builder /whitespace/wsc /whitespace/ws2c /whitespace/wsa \
                    /whitespace/gen*.c /whitespace/io.h /whitespace/memory.h \
                    /whitespace/*.ws /
COPY whitespace/tests /tests
ENTRYPOINT ["/ws2c"]
