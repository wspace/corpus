FROM golang:1.22 AS builder

WORKDIR /wspacego
COPY wspacego .
RUN go mod init github.com/135yshr/wspacego && \
    go mod tidy && \
    CGO_ENABLED=0 go build
# Tests fail
# RUN go test ./...

FROM scratch

COPY --from=builder /wspacego/wspacego /
ENTRYPOINT ["/wspacego"]
