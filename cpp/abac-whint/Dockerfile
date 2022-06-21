FROM alpine AS builder

RUN apk add g++
WORKDIR /WHINT
COPY WHINT .
RUN g++ -O3 -Wall -static -o whint whint.cpp

FROM scratch

COPY --from=builder /WHINT/whint /
ENTRYPOINT ["/whint"]
