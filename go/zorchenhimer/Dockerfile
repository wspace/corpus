FROM golang:1.22 AS builder

WORKDIR /whitespace
COPY whitespace .
RUN go test . && \
    CGO_ENABLED=0 go build -o bin/wt cmd/wt.go && \
    CGO_ENABLED=0 go build -o bin/wi cmd/wi.go

FROM scratch

COPY --from=builder /whitespace/bin/wt /whitespace/bin/wi /
COPY whitespace/code-examples /
ENTRYPOINT ["/wi"]
