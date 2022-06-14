FROM golang:1.18 AS builder

WORKDIR /
RUN git clone https://github.com/wspace/technohippy-go go-whitespace
WORKDIR /go-whitespace
RUN go build -o go-whitespace ./src

FROM scratch

COPY --from=builder /go-whitespace/go-whitespace /
ENTRYPOINT ["/go-whitespace"]
