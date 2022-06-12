FROM alpine as builder

RUN apk add git make g++
RUN git clone https://github.com/benajmin/whitespace-interpreter
WORKDIR /whitespace-interpreter
RUN make

FROM scratch as runner

COPY --from=builder /whitespace-interpreter/WhitespaceInterpreter.out /
ENTRYPOINT ["/WhitespaceInterpreter.out"]
