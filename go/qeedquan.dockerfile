FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/wspace/qeedquan-go whitespace
WORKDIR /home/whitespace
RUN go mod init github.com/wspace/qeedquan-go
RUN go build -o whitespace
# builds: /home/whitespace/whitespace
