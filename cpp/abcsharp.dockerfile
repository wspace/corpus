FROM alpine AS builder

RUN apk add git g++
RUN git clone https://github.com/abcsharp/Whitespace
WORKDIR /Whitespace
RUN g++ -O3 -Wall -std=c++11 -static -o wsi whitespace/main.cpp

FROM scratch

COPY --from=builder /Whitespace/wsi /
ENTRYPOINT ["/wsi"]
