FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/wspace/qeedquan-go
WORKDIR /home/qeedquan-go
RUN go mod init github.com/wspace/qeedquan-go
RUN go build -o whitespace
# builds: /home/qeedquan-go/whitespace
