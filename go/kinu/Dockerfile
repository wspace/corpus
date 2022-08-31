FROM golang:1.19 AS builder

WORKDIR /whitespace
COPY whitespace .
RUN go mod init github.com/kinu/whitespace
RUN go build -buildvcs=false

FROM scratch

COPY --from=builder /whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
