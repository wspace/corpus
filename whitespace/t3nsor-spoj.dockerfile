FROM alpine as builder

RUN apk add git g++
RUN git clone https://github.com/t3nsor/SPOJ
WORKDIR /SPOJ
RUN g++ -g -std=c++17 -o $0 wstp.cpp

FROM scratch as runner

COPY --from=builder /SPOJ/wstp /
ENTRYPOINT ["/wstp"]
