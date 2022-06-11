FROM alpine

RUN apk add git make gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/rispoli/txt2ws
WORKDIR /home/txt2ws
RUN make
RUN test -f /home/txt2ws/txt2ws
