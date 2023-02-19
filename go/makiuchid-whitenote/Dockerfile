FROM golang:1.20 AS builder

WORKDIR /
COPY whitenote /whitenote
WORKDIR /whitenote/wspace/cmd/wspace
RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder /whitenote/wspace/cmd/wspace/wspace /
ENTRYPOINT ["/wspace"]
