FROM alpine AS builder

RUN apk add git make gcc musl-dev
RUN git clone https://github.com/wspace/stellwag-wspacegen wspacegen
WORKDIR /wspacegen
RUN make

FROM scratch

COPY --from=builder /wspacegen/wspacegen /
COPY --from=builder /wspacegen/wsdbg /
COPY --from=builder /wspacegen/wsi /
ENTRYPOINT ["/wspacegen"]
