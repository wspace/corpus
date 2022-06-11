FROM alpine

RUN apk add git make g++
WORKDIR /home
RUN git clone https://github.com/benajmin/whitespace-interpreter
WORKDIR /home/whitespace-interpreter
RUN make
RUN test -f /home/whitespace-interpreter/WhitespaceInterpreter.out
