FROM alpine AS builder

RUN apk add g++
WORKDIR /SPOJ
COPY SPOJ .
RUN g++ -O3 -static -o wstp wstp.cpp

FROM scratch

COPY --from=builder /SPOJ/wstp /
COPY SPOJ/sbstr1 SPOJ/sbstr1.ws /
ENTRYPOINT ["/wstp"]
