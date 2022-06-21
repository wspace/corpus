FROM golang:1.18 AS builder

WORKDIR /
RUN git clone https://github.com/samuel-pratt/whitespace-go
WORKDIR /whitespace-go
RUN go build

FROM scratch

COPY --from=builder /whitespace-go/whitespace-go /
ENTRYPOINT ["/whitespace-go"]
