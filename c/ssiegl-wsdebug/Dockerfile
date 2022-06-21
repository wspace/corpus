FROM alpine AS builder

RUN apk add make gcc g++ gmp-dev readline-dev
WORKDIR /wsdebug
COPY wsdebug .
RUN ./configure && make

FROM scratch

COPY --from=builder /wsdebug/wsdebug /
COPY --from=builder /wsdebug/wsi /
ENTRYPOINT ["/wsdebug"]
