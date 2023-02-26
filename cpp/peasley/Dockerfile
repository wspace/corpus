FROM alpine AS builder

RUN apk add --no-cache g++
WORKDIR /whitespace
COPY whitespace .
RUN g++ -O3 -static -o whitespace whitespace.cpp

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
