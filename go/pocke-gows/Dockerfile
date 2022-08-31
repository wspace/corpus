FROM golang:1.19 AS builder

WORKDIR /gows
COPY gows .
RUN go mod init github.com/pocke/gows && \
    go mod tidy
RUN go build -buildvcs=false

FROM scratch

COPY --from=builder /gows/gows /
ENTRYPOINT ["/gows"]
