FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/wspace/technohippy-go
WORKDIR /home/technohippy-go
RUN go build -o go-whitespace ./src
# builds: /home/technohippy-go/go-whitespace
