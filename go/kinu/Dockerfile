FROM golang:1.18 AS builder

WORKDIR /
WORKDIR /whitespace
COPY whitespace .
RUN go mod init github.com/kinu/whitespace
RUN go build

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
