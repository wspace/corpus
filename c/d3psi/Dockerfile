FROM alpine AS builder

RUN apk add --no-cache cmake make gcc g++
WORKDIR /whitespace-interpreter
COPY whitespace-interpreter/src .
RUN gcc -O3 -static -o interpreter *.c

FROM scratch

COPY --from=builder /whitespace-interpreter/interpreter /
COPY whitespace-interpreter/ws/*.ws /
ENTRYPOINT ["/interpreter"]
