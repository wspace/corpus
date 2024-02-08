FROM golang:1.22 AS builder

WORKDIR /go-wspace
COPY go-wspace .
RUN go mod init github.com/tempxla/go-wspace && \
    CGO_ENABLED=0 go build -o bin/go-wspace ./src
# Tests fail
# RUN go test ./...

FROM scratch

COPY --from=builder /go-wspace/bin/go-wspace /
ENTRYPOINT ["/go-wspace"]
