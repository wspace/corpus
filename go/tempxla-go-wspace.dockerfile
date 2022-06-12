FROM golang:1.18 as builder

WORKDIR /
RUN git clone https://github.com/tempxla/go-wspace
WORKDIR /go-wspace
RUN go mod init github.com/tempxla/go-wspace
# RUN go test ./...
RUN go build -o bin/go-wspace ./src

FROM scratch as runner

COPY --from=builder /go-wspace/bin/go-wspace /
ENTRYPOINT ["/go-wspace"]
