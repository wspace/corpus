FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/tempxla/go-wspace
WORKDIR /home/go-wspace
RUN go mod init github.com/tempxla/go-wspace
RUN go build -o bin/go-wspace ./src
# builds: /home/go-wspace/bin/go-wspace
