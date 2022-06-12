FROM alpine as builder

RUN apk add git g++
RUN git clone https://github.com/wspace/peasley-cpp whitespace
WORKDIR /whitespace
RUN g++ -O3 -Wall -o whitespace whitespace.cpp

FROM scratch as runner

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
