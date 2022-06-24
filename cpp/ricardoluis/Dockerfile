FROM alpine AS builder

RUN apk add g++
WORKDIR /whitespace
COPY whitespace .
RUN g++ -O3 -std=c++11 -static -o whitespace whitespace.cpp

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
