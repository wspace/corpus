FROM alpine as builder

RUN apk add git make g++ gmp-dev
RUN git clone https://git.code.sf.net/p/esco/code esco
WORKDIR /esco
RUN ./configure && make

FROM scratch as runner

COPY --from=builder /esco/src/esco /
ENTRYPOINT ["/esco"]
