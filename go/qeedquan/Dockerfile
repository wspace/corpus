FROM golang:1.18 AS builder

WORKDIR /
WORKDIR /whitespace
COPY whitespace .
RUN go mod init github.com/wspace/qeedquan-go
RUN go build -o whitespace

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
