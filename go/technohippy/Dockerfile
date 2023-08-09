FROM golang:1.21 AS builder

WORKDIR /go-whitespace
COPY go-whitespace .
RUN CGO_ENABLED=0 go build -o go-whitespace ./src

FROM scratch

COPY --from=builder /go-whitespace/go-whitespace /
ENTRYPOINT ["/go-whitespace"]
