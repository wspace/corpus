FROM golang:1.18 as builder

WORKDIR /
RUN git clone https://github.com/wspace/qeedquan-go whitespace
WORKDIR /whitespace
RUN go mod init github.com/wspace/qeedquan-go
RUN go build -o whitespace

FROM scratch as runner

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
