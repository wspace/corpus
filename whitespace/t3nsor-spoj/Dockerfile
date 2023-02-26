FROM alpine AS builder

RUN apk add --no-cache g++
WORKDIR /SPOJ
COPY SPOJ/wstp.cpp .
RUN g++ -O3 -static -o wstp wstp.cpp

FROM scratch

COPY --from=builder /SPOJ/wstp /
COPY SPOJ/sbstr* /
ENTRYPOINT ["/wstp"]
