FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /whitespace
COPY whitespace .
RUN gcc -O3 -static -o whitespace whitespace.c

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
