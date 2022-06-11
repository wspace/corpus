FROM alpine

RUN apk add git make gcc g++
WORKDIR /home
RUN git clone https://github.com/wspace/ssiegl-wsdebug wsdebug
WORKDIR /home/wsdebug
RUN ./configure
RUN make
RUN test -f /home/wsdebug/wsdebug
RUN test -f /home/wsdebug/wsi
