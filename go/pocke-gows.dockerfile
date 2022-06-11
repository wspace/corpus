FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/pocke/gows
WORKDIR /home/gows
RUN go mod init github.com/pocke/gows
RUN go mod tidy
RUN go build
RUN test -f /home/gows/gows
