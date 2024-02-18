FROM golang:1.22 AS builder

WORKDIR /whitespace
COPY whitespace .
RUN go mod init github.com/wspace/qeedquan-go && \
    CGO_ENABLED=0 go build -o whitespace

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
