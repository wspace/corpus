FROM golang:1.22 AS builder

WORKDIR /whitespace_go
COPY whitespace_go .
RUN go test ./... && \
    CGO_ENABLED=0 go build -o releases/ws cmd/ws.go

FROM scratch

COPY --from=builder /whitespace_go/releases/ws /
ENTRYPOINT ["/ws"]
