FROM golang:1.18 AS builder

WORKDIR /
WORKDIR /gows
COPY gows .
RUN go mod init github.com/pocke/gows
RUN go mod tidy
RUN go build

FROM scratch

COPY --from=builder /gows/gows /
ENTRYPOINT ["/gows"]
