FROM golang:1.18 as builder

WORKDIR /
RUN git clone https://github.com/samuel-pratt/whitespace-go
WORKDIR /whitespace-go
RUN go build

FROM scratch as runner

COPY --from=builder /whitespace-go/whitespace-go /
ENTRYPOINT ["/whitespace-go"]
