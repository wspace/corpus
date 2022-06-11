FROM alpine

RUN apk add git make gcc musl-dev flex bison
WORKDIR /home
RUN git clone https://gitlab.com/tejaskasetty/ws-compiler
WORKDIR /home/ws-compiler/flex-bison/262_267_256
RUN make
RUN test -f /home/ws-compiler/flex-bison/262_267_256/ws
