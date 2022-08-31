FROM golang:1.19 AS builder

WORKDIR /go-wspace
COPY go-wspace .
RUN go mod init github.com/tempxla/go-wspace
# RUN go test ./...
RUN go build -buildvcs=false -o bin/go-wspace ./src

FROM scratch

COPY --from=builder /go-wspace/bin/go-wspace /
ENTRYPOINT ["/go-wspace"]
