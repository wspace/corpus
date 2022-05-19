FROM golang:1.18

RUN git clone https://github.com/simomu-github/whitespace_go /simomu-go
WORKDIR /simomu-go
RUN go build -o releases/ws cmd/ws.go
