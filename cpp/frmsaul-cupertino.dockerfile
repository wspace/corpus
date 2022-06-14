FROM alpine AS builder

RUN apk add git g++
RUN git clone https://github.com/frmsaul/Cupertino-WhiteSpace-Interperter
WORKDIR /Cupertino-WhiteSpace-Interperter
RUN g++ -O3 -static -o whitespace src/*.cpp

FROM scratch

COPY --from=builder /Cupertino-WhiteSpace-Interperter/whitespace /
ENTRYPOINT ["/whitespace"]
