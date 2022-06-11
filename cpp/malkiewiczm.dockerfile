FROM alpine

RUN apk add git make g++ flex bison
WORKDIR /home
RUN git clone https://github.com/malkiewiczm/whitespace_compiler
WORKDIR /home/whitespace_compiler
RUN make
RUN test -f /home/whitespace_compiler/compile
