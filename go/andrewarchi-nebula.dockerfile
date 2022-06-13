FROM golang:1.18 as builder

WORKDIR /
RUN git clone https://github.com/andrewarchi/nebula
WORKDIR /nebula
# Requires LLVM
RUN go mod init github.com/andrewarchi/nebula
RUN go mod tidy
RUN go build

FROM scratch as runner

COPY --from=builder /nebula/nebula /
ENTRYPOINT ["/nebula"]
