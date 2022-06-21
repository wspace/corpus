FROM alpine AS builder

RUN apk add make gcc g++
WORKDIR /wsdebug
COPY wsdebug .
RUN ./configure
RUN make

FROM scratch

COPY --from=builder /wsdebug/wsdebug /
COPY --from=builder /wsdebug/wsi /
ENTRYPOINT ["/wsdebug"]
