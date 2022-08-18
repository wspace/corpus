FROM golang:1.19 AS builder

WORKDIR /
COPY whitenote /whitenote
WORKDIR /whitenote/wspace/cmd/wspace
RUN go build -buildvcs=false

FROM scratch

COPY --from=builder /whitenote/wspace/cmd/wspace/wspace /
ENTRYPOINT ["/wspace"]
