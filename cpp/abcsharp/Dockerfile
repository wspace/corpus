FROM alpine AS builder

RUN apk add --no-cache g++
WORKDIR /Whitespace
COPY Whitespace .
RUN g++ -O3 -std=c++11 -static -o wsi whitespace/main.cpp

FROM scratch

COPY --from=builder /Whitespace/wsi /
ENTRYPOINT ["/wsi"]
