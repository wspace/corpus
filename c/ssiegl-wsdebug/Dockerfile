FROM alpine AS builder

WORKDIR /wsdebug
RUN <<EOF
set -e

apk add --no-cache curl make gcc g++ gmp-dev readline-dev
# wsdebug-0.1.tar.gz has ./configure, but not git
curl https://master.dl.sourceforge.net/project/wsdebug.berlios/wsdebug-0.1.tar.gz?viasf=1 | gunzip | tar xf - --strip-components 1
./configure
make CFLAGS='-O3 -g -static'

EOF

FROM alpine

RUN apk add --no-cache gmp readline
COPY --from=builder /wsdebug/wsdebug /wsdebug/wsi /
ENTRYPOINT ["/wsdebug"]
