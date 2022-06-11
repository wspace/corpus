FROM alpine

RUN apk add git make gcc musl-dev flex bison
WORKDIR /home
RUN git clone https://github.com/wspace/remigascou-c whitespace
WORKDIR /home/whitespace/dev
RUN make
WORKDIR /home/whitespace/lexxyacc
RUN make
RUN test -f /home/whitespace/dev/bin/compiler
RUN test -f /home/whitespace/dev/bin/decompiler
RUN test -f /home/whitespace/lexxyacc/compiler/compiler
