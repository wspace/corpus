FROM alpine AS builder

RUN apk add make g++
WORKDIR /ws-interpreter
COPY ws-interpreter .
RUN mkdir bin
RUN make -C src TARGET=../bin/ws CXX='g++ -static'

FROM scratch

COPY --from=builder /ws-interpreter/bin/ws /
ENTRYPOINT ["/ws"]
