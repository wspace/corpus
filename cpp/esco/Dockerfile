FROM alpine AS builder

RUN apk add --no-cache make g++ gmp-dev
WORKDIR /esco
COPY esco .
RUN ./configure && make

FROM scratch

COPY --from=builder /esco/src/esco /
ENTRYPOINT ["/esco"]
