FROM golang:1.22 AS builder

WORKDIR /nebula
COPY nebula .
# TODO: Requires LLVM
RUN go mod init github.com/thaliaarchi/nebula && \
    go mod tidy && \
    go test ./... && \
    go build

FROM scratch

COPY --from=builder /nebula/nebula /
ENTRYPOINT ["/nebula"]
