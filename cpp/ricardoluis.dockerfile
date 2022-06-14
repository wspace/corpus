FROM alpine AS builder

RUN apk add git g++
RUN git clone https://github.com/RicardoLuis0/whitespace
WORKDIR /whitespace
RUN g++ -O3 -Wall -std=c++11 -static -o whitespace whitespace.cpp

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
