FROM alpine as builder

RUN apk add git g++
RUN git clone https://github.com/abcsharp/Whitespace
WORKDIR /Whitespace
RUN g++ -O3 -Wall -std=c++11 -o wsi whitespace/main.cpp

FROM scratch as runner

COPY --from=builder /Whitespace/wsi /
ENTRYPOINT ["/wsi"]
