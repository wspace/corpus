FROM alpine AS builder

RUN apk add --no-cache curl make gcc g++ gmp-dev readline-dev
WORKDIR /wsdebug
# wsdebug-0.1.tar.gz has ./configure, but not git
RUN curl https://master.dl.sourceforge.net/project/wsdebug.berlios/wsdebug-0.1.tar.gz?viasf=1 | gunzip | tar xf - --strip-components 1
RUN ./configure && \
    make CFLAGS='-O3 -g -static'

FROM alpine

RUN apk add --no-cache gmp readline
COPY --from=builder /wsdebug/wsdebug /wsdebug/wsi /
ENTRYPOINT ["/wsdebug"]
