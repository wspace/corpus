FROM alpine AS builder

RUN apk add --no-cache g++
WORKDIR /whitespace
COPY whitespace .
RUN g++ -O3 -static -o white main.cpp

FROM scratch

COPY --from=builder /whitespace/white /
ENTRYPOINT ["/white"]
