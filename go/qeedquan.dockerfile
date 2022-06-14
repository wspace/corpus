FROM golang:1.18 AS builder

WORKDIR /
RUN git clone https://github.com/wspace/qeedquan-go whitespace
WORKDIR /whitespace
RUN go mod init github.com/wspace/qeedquan-go
RUN go build -o whitespace

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
