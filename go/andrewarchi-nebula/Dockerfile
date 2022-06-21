FROM golang:1.18 AS builder

WORKDIR /
WORKDIR /nebula
COPY nebula .
# Requires LLVM
RUN go mod init github.com/andrewarchi/nebula
RUN go mod tidy
RUN go build

FROM scratch

COPY --from=builder /nebula/nebula /
ENTRYPOINT ["/nebula"]
