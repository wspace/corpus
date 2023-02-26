FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /whitespace-php
COPY whitespace-php .
RUN gcc -O3 -static -o whitespace *.c

FROM scratch

COPY --from=builder /whitespace-php/whitespace /
COPY whitespace-php/samples /
ENTRYPOINT ["/whitespace"]
