FROM golang:1.18

RUN git clone https://github.com/wspace/qeedquan-go /qeedquan-go
WORKDIR /qeedquan-go
RUN go mod init github.com/wspace/qeedquan-go
RUN go build -o whitespace
