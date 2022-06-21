FROM alpine AS builder

RUN apk add g++
WORKDIR /SPOJ
COPY SPOJ .
RUN g++ -g -std=c++17 -o $0 wstp.cpp

FROM scratch

COPY --from=builder /SPOJ/wstp /
ENTRYPOINT ["/wstp"]
