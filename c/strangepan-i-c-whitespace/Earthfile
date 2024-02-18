FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /I_C_Whitespace
COPY I_C_Whitespace .
RUN gcc -O3 -g -static -std=c99 -o whitespace *.c

FROM scratch

COPY --from=builder /I_C_Whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
