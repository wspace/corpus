FROM alpine as builder

RUN apk add git g++
RUN git clone https://github.com/benajmin/whitespace-interpreter
WORKDIR /whitespace-interpreter
RUN g++ -Wall -static -o WhitespaceInterpreter.out *.cpp

FROM scratch as runner

COPY --from=builder /whitespace-interpreter/WhitespaceInterpreter.out /
ENTRYPOINT ["/WhitespaceInterpreter.out"]
