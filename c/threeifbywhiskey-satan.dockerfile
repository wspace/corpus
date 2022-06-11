FROM alpine

RUN apk add git make gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/threeifbywhiskey/satan
WORKDIR /home/satan
RUN make
RUN test -f /home/satan/satan
