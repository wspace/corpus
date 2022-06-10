FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/kinu/whitespace
WORKDIR /home/whitespace
RUN go mod init github.com/kinu/whitespace
RUN go build
# builds: /home/whitespace/whitespace
