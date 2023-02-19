FROM golang:1.20 AS builder

WORKDIR /whitespace-go
COPY whitespace-go .
RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder /whitespace-go/whitespace-go /
ENTRYPOINT ["/whitespace-go"]
