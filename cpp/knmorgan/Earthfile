FROM alpine AS builder

RUN apk add --no-cache make g++ gmp-dev
WORKDIR /ws
COPY ws .
RUN make CC='g++ -static'

FROM scratch

COPY --from=builder /ws/ws /
ENTRYPOINT ["/ws"]
