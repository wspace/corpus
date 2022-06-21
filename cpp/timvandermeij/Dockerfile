FROM alpine AS builder

RUN apk add make g++
WORKDIR /whitespace-interpreter
COPY whitespace-interpreter .
RUN make

FROM scratch

COPY --from=builder /whitespace-interpreter/whitespace /
ENTRYPOINT ["/whitespace"]
