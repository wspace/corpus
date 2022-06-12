FROM alpine as builder

RUN apk add git make gcc g++
RUN git clone https://github.com/wspace/ssiegl-wsdebug wsdebug
WORKDIR /wsdebug
RUN ./configure
RUN make

FROM scratch as runner

COPY --from=builder /wsdebug/wsdebug /
COPY --from=builder /wsdebug/wsi /
ENTRYPOINT ["/wsdebug"]
