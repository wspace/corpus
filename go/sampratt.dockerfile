FROM golang:1.18

RUN git clone https://github.com/samuel-pratt/whitespace-go /sampratt-go
WORKDIR /sampratt-go
RUN go build
# builds: /sampratt-go/whitespace-go
