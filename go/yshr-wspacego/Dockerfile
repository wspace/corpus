FROM golang:1.20 AS builder

WORKDIR /wspacego
COPY wspacego .
RUN go mod init github.com/135yshr/wspacego && \
    go mod tidy
# RUN go test ./...
RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder /wspacego/wspacego /
ENTRYPOINT ["/wspacego"]
