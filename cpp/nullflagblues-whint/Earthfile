FROM alpine AS builder

RUN apk add --no-cache g++
WORKDIR /WHINT
COPY WHINT .
RUN g++ -O3 -static -o whint whint.cpp

FROM scratch

COPY --from=builder /WHINT/whint /
ENTRYPOINT ["/whint"]
