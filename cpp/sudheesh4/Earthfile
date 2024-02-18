FROM alpine AS builder

RUN apk add --no-cache g++
WORKDIR /Whitespace-Interpreter-C-
COPY Whitespace-Interpreter-C- .
RUN g++ -O3 -static -o space space.cpp

FROM scratch

COPY --from=builder /Whitespace-Interpreter-C-/space /
ENTRYPOINT ["/space"]
