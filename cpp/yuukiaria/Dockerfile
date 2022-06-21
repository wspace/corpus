FROM alpine AS builder

RUN apk add git make g++
RUN git clone https://github.com/YuukiARIA/ws-interpreter
WORKDIR /ws-interpreter/src
RUN mkdir ../bin
RUN make TARGET=../bin/ws CXX='g++ -static'

FROM scratch

COPY --from=builder /ws-interpreter/bin/ws /
ENTRYPOINT ["/ws"]
