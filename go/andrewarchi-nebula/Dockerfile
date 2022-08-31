FROM golang:1.19 AS builder

WORKDIR /nebula
COPY nebula .
# TODO: Requires LLVM
RUN go mod init github.com/andrewarchi/nebula && \
    go mod tidy
RUN go test ./...
RUN go build -buildvcs=false

FROM scratch

COPY --from=builder /nebula/nebula /
ENTRYPOINT ["/nebula"]
