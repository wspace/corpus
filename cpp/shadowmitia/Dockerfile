FROM alpine AS builder

RUN apk add bash g++
WORKDIR /whitespace
COPY whitespace .
RUN bash build.sh release wspace

FROM scratch

COPY --from=builder /whitespace/wspace /
ENTRYPOINT ["/wspace"]
