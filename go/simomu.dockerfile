FROM golang:1.18

WORKDIR /home
RUN git clone https://github.com/simomu-github/whitespace_go
WORKDIR /home/whitespace_go
RUN go build -o releases/ws cmd/ws.go
# builds: /home/whitespace_go/releases/ws
