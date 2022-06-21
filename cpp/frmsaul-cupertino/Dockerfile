FROM alpine AS builder

RUN apk add g++
WORKDIR /Cupertino-WhiteSpace-Interperter
COPY Cupertino-WhiteSpace-Interperter .
RUN g++ -O3 -static -o whitespace src/*.cpp

FROM scratch

COPY --from=builder /Cupertino-WhiteSpace-Interperter/whitespace /
ENTRYPOINT ["/whitespace"]
