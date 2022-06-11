FROM alpine

RUN apk add git make g++ gmp-dev
WORKDIR /home
RUN git clone https://github.com/knmorgan/ws
WORKDIR /home/ws
RUN make CXXFLAGS='-O3 -Wall -pedantic'
RUN test -f /home/ws/ws
