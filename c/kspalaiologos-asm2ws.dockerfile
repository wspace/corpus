FROM alpine

RUN apk add git make gcc musl-dev flex bison
WORKDIR /home
RUN git clone https://github.com/kspalaiologos/asm2ws
WORKDIR /home/asm2ws
RUN ./configure --with-target=release
RUN make -j4 wsi
RUN test -f /home/asm2ws/wsi
