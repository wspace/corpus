FROM golang:1.18 as builder

WORKDIR /
RUN git clone https://github.com/wspace/technohippy-go go-whitespace
WORKDIR /go-whitespace
RUN go build -o go-whitespace ./src

FROM scratch as runner

COPY --from=builder /go-whitespace/go-whitespace /
ENTRYPOINT ["/go-whitespace"]
