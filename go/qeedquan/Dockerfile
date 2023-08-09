FROM golang:1.21 AS builder

WORKDIR /whitespace
COPY whitespace .
RUN go mod init github.com/wspace/qeedquan-go
RUN CGO_ENABLED=0 go build -o whitespace

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
