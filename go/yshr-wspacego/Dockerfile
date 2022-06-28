FROM golang:1.18 AS builder

WORKDIR /
WORKDIR /wspacego
COPY wspacego .
RUN go mod init github.com/135yshr/wspacego && \
    go mod tidy
# RUN go test ./...
RUN go build

FROM scratch

COPY --from=builder /wspacego/wspacego /
ENTRYPOINT ["/wspacego"]
