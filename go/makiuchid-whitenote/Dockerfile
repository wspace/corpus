FROM golang:1.22 AS builder

WORKDIR /whitenote
COPY whitenote .
WORKDIR /whitenote/wspace/cmd/wspace
RUN go test && \
    CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder /whitenote/wspace/cmd/wspace/wspace /
ENTRYPOINT ["/wspace"]
