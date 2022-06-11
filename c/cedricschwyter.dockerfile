FROM alpine

RUN apk add git cmake make gcc g++
WORKDIR /home
RUN git clone https://github.com/D3PSI/whitespace-interpreter
WORKDIR /home/whitespace-interpreter
RUN cmake .
RUN make
RUN test -f /home/whitespace-interpreter/interpreter
