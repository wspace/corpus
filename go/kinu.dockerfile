FROM golang:1.18

RUN git clone https://github.com/kinu/whitespace /kinu-go
WORKDIR /kinu-go
RUN go mod init github.com/kinu/whitespace
RUN go build
# builds: /kinu-go/whitespace
