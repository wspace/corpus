FROM alpine AS builder

RUN apk add --no-cache make crystal
WORKDIR /spiceweight
COPY spiceweight .
RUN make

FROM scratch

COPY --from=builder /spiceweight/spwt /
ENTRYPOINT ["/spwt"]
