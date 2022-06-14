FROM alpine AS builder

RUN apk add git g++
RUN git clone https://github.com/wspace/peasley-cpp whitespace
WORKDIR /whitespace
RUN g++ -O3 -Wall -static -o whitespace whitespace.cpp

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
