FROM alpine

RUN apk add git make g++
WORKDIR /home
RUN git clone https://github.com/wspace/marcellippmann-whitepp Whitepp
WORKDIR /home/Whitepp
RUN make
RUN test -f /home/Whitepp/bin/White++
