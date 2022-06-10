FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/wspace/technohippy-go go-whitespace
WORKDIR /home/go-whitespace
RUN go build -o go-whitespace ./src
# builds: /home/go-whitespace/go-whitespace
