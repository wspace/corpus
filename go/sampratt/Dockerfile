FROM golang:1.19 AS builder

WORKDIR /whitespace-go
COPY whitespace-go .
RUN go build -buildvcs=false

FROM scratch

COPY --from=builder /whitespace-go/whitespace-go /
ENTRYPOINT ["/whitespace-go"]
