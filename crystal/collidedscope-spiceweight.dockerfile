FROM alpine

RUN apk add git make crystal
WORKDIR /home
RUN git clone https://github.com/collidedscope/spiceweight
WORKDIR /home/spiceweight
RUN make
RUN test -f /home/spiceweight/spwt
