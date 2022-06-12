FROM alpine as builder

RUN apk add git g++
RUN git clone https://github.com/Keirua/whitespace
WORKDIR /whitespace
RUN g++ -O3 -Wall -o white main.cpp

FROM scratch as runner

COPY --from=builder /whitespace/white /
ENTRYPOINT ["/white"]
