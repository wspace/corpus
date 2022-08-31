FROM golang:1.19 AS builder

WORKDIR /go-whitespace
COPY go-whitespace .
RUN go build -buildvcs=false -o go-whitespace ./src

FROM scratch

COPY --from=builder /go-whitespace/go-whitespace /
ENTRYPOINT ["/go-whitespace"]
