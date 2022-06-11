FROM alpine

RUN apk add git make gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/threeifbywhiskey/blacktime
WORKDIR /home/blacktime
RUN make
RUN test -f /home/blacktime/blacktime
