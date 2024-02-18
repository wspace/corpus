FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /whitespace-interpreter
COPY whitespace-interpreter/src .
RUN gcc -O3 -g -static -o interpreter *.c

FROM scratch

COPY --from=builder /whitespace-interpreter/interpreter /
COPY whitespace-interpreter/ws/*.ws /
ENTRYPOINT ["/interpreter"]
