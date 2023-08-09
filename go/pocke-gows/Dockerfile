FROM golang:1.21 AS builder

WORKDIR /gows
COPY gows .
RUN go mod init github.com/pocke/gows && \
    go mod tidy
RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder /gows/gows /
ENTRYPOINT ["/gows"]
