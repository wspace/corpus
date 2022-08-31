FROM golang:1.19 AS builder

WORKDIR /whitespace
COPY whitespace .
RUN go mod init github.com/wspace/qeedquan-go
RUN go build -buildvcs=false -o whitespace

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
