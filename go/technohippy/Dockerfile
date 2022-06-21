FROM golang:1.18 AS builder

WORKDIR /
WORKDIR /go-whitespace
COPY go-whitespace .
RUN go build -o go-whitespace ./src

FROM scratch

COPY --from=builder /go-whitespace/go-whitespace /
ENTRYPOINT ["/go-whitespace"]
