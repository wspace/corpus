FROM alpine AS builder

RUN apk add git bash g++
RUN git clone https://github.com/ShadowMitia/whitespace
WORKDIR /whitespace
RUN bash build.sh release wspace

FROM scratch

COPY --from=builder /whitespace/wspace /
ENTRYPOINT ["/wspace"]
