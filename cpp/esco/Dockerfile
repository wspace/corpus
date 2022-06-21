FROM alpine AS builder

RUN apk add git make g++ gmp-dev
RUN git clone https://git.code.sf.net/p/esco/code esco
WORKDIR /esco
RUN ./configure && make

FROM scratch

COPY --from=builder /esco/src/esco /
ENTRYPOINT ["/esco"]
