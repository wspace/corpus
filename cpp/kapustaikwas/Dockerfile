FROM alpine AS builder

RUN apk add git g++
RUN git clone https://github.com/kapustaikwas27/Whitespace-compiler
WORKDIR /Whitespace-compiler
RUN g++ -O3 -Wall -static -o pre pre.cpp

FROM scratch

COPY --from=builder /Whitespace-compiler/pre /
ENTRYPOINT ["/pre"]
