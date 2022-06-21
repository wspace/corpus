FROM golang:1.18 AS builder

WORKDIR /
RUN git clone https://github.com/tempxla/go-wspace
WORKDIR /go-wspace
RUN go mod init github.com/tempxla/go-wspace
# RUN go test ./...
RUN go build -o bin/go-wspace ./src

FROM scratch

COPY --from=builder /go-wspace/bin/go-wspace /
ENTRYPOINT ["/go-wspace"]
