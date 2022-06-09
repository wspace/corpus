FROM golang:1.18

RUN git clone https://github.com/tempxla/go-wspace /tempxla-go-wspace
WORKDIR /tempxla-go-wspace
RUN go mod init github.com/tempxla/go-wspace
RUN go build -o bin/go-wspace ./src
# builds: /tempxla-go-wspace/bin/go-wspace
