FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev
WORKDIR /wspacegen
COPY wspacegen .
RUN make CFLAGS='-O3 -g -static'

FROM scratch

COPY --from=builder /wspacegen/wspacegen /wspacegen/wsdbg /wspacegen/wsi /
ENTRYPOINT ["/wspacegen"]
