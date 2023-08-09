FROM golang:1.21 AS builder

WORKDIR /whitespace_go
COPY whitespace_go .
RUN go test ./...
RUN CGO_ENABLED=0 go build -o releases/ws cmd/ws.go

FROM scratch

COPY --from=builder /whitespace_go/releases/ws /
ENTRYPOINT ["/ws"]
