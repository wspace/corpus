FROM alpine AS builder

RUN apk add g++
WORKDIR /whitespace
COPY whitespace .
RUN g++ -O3 -Wall -static -o white main.cpp

FROM scratch

COPY --from=builder /whitespace/white /
ENTRYPOINT ["/white"]
