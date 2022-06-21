FROM golang:1.18 AS builder

WORKDIR /
RUN git clone https://github.com/135yshr/wspacego
WORKDIR /wspacego
RUN go mod init github.com/135yshr/wspacego
RUN go mod tidy
# RUN go test ./...
RUN go build

FROM scratch

COPY --from=builder /wspacego/wspacego /
ENTRYPOINT ["/wspacego"]
