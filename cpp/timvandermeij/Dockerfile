FROM alpine AS builder

RUN apk add --no-cache make g++
WORKDIR /whitespace-interpreter
COPY whitespace-interpreter .
RUN make

FROM scratch

COPY --from=builder /whitespace-interpreter/whitespace /
ENTRYPOINT ["/whitespace"]
