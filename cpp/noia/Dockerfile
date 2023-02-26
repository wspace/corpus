FROM alpine AS builder

RUN apk add --no-cache g++
WORKDIR /Whitespace-Interpreter
COPY Whitespace-Interpreter .
RUN g++ -O3 -static -o ws Lexer.cc Parser.cc

FROM scratch

COPY --from=builder /Whitespace-Interpreter/ws /
ENTRYPOINT ["/ws"]
