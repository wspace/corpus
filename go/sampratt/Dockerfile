FROM golang:1.18 AS builder

WORKDIR /
WORKDIR /whitespace-go
COPY whitespace-go .
RUN go build

FROM scratch

COPY --from=builder /whitespace-go/whitespace-go /
ENTRYPOINT ["/whitespace-go"]
