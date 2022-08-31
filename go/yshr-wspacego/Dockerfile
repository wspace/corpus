FROM golang:1.19 AS builder

WORKDIR /wspacego
COPY wspacego .
RUN go mod init github.com/135yshr/wspacego && \
    go mod tidy
# RUN go test ./...
RUN go build -buildvcs=false

FROM scratch

COPY --from=builder /wspacego/wspacego /
ENTRYPOINT ["/wspacego"]
