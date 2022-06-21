FROM alpine AS builder

RUN apk add git g++
RUN git clone https://github.com/abac00s/WHINT
WORKDIR /WHINT
RUN g++ -O3 -Wall -static -o whint whint.cpp

FROM scratch

COPY --from=builder /WHINT/whint /
ENTRYPOINT ["/whint"]
