FROM alpine

RUN apk add git make gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/wspace/stellwag-wspacegen wspacegen
WORKDIR /home/wspacegen
RUN make
RUN test -f /home/wspacegen/wspacegen
RUN test -f /home/wspacegen/wsdbg
RUN test -f /home/wspacegen/wsi
