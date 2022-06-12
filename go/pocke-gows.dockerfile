FROM golang:1.18 as builder

WORKDIR /
RUN git clone https://github.com/pocke/gows
WORKDIR /gows
RUN go mod init github.com/pocke/gows
RUN go mod tidy
RUN go build

FROM scratch as runner

COPY --from=builder /gows/gows /
ENTRYPOINT ["/gows"]
