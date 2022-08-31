FROM golang:1.19 AS builder

WORKDIR /whitespace_go
COPY whitespace_go .
RUN go test ./...
RUN go build -buildvcs=false -o releases/ws cmd/ws.go

FROM scratch

COPY --from=builder /whitespace_go/releases/ws /
ENTRYPOINT ["/ws"]
