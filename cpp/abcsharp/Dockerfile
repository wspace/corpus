FROM alpine AS builder

RUN apk add g++
WORKDIR /Whitespace
COPY Whitespace .
RUN g++ -O3 -Wall -std=c++11 -static -o wsi whitespace/main.cpp

FROM scratch

COPY --from=builder /Whitespace/wsi /
ENTRYPOINT ["/wsi"]
