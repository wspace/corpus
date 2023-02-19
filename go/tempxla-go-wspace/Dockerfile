FROM golang:1.20 AS builder

WORKDIR /go-wspace
COPY go-wspace .
RUN go mod init github.com/tempxla/go-wspace
# RUN go test ./...
RUN CGO_ENABLED=0 go build -o bin/go-wspace ./src

FROM scratch

COPY --from=builder /go-wspace/bin/go-wspace /
ENTRYPOINT ["/go-wspace"]
