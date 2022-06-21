FROM alpine AS builder

RUN apk add make g++
WORKDIR /respace
COPY respace .
RUN make test
RUN make LDFLAGS=-static

FROM scratch

COPY --from=builder /respace/respace /
ENTRYPOINT ["/respace"]
