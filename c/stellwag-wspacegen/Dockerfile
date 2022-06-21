FROM alpine AS builder

RUN apk add make gcc musl-dev
WORKDIR /wspacegen
COPY wspacegen .
RUN make

FROM scratch

COPY --from=builder /wspacegen/wspacegen /
COPY --from=builder /wspacegen/wsdbg /
COPY --from=builder /wspacegen/wsi /
ENTRYPOINT ["/wspacegen"]
