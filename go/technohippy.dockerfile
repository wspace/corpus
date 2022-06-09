FROM golang:1.18

RUN git clone https://github.com/wspace/technohippy-go /technohippy-go
WORKDIR /technohippy-go
RUN go build -o go-whitespace ./src
# builds: /technohippy-go/go-whitespace
