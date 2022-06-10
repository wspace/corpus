FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/samuel-pratt/whitespace-go
WORKDIR /home/whitespace-go
RUN go build
# builds: /home/whitespace-go/whitespace-go
