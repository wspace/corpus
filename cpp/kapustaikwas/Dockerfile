FROM alpine AS builder

RUN apk add --no-cache g++
WORKDIR /Whitespace-compiler
COPY Whitespace-compiler .
RUN g++ -O3 -static -o pre pre.cpp

FROM scratch

COPY --from=builder /Whitespace-compiler/pre /
ENTRYPOINT ["/pre"]
