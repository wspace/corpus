FROM golang:1.18

RUN git clone https://github.com/pocke/gows /pocke-gows
WORKDIR /pocke-gows
RUN go mod init github.com/pocke/gows
RUN go mod tidy
RUN go build
