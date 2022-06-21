FROM alpine AS builder

RUN apk add make g++ gmp-dev
WORKDIR /ws
COPY ws .
RUN make CC='g++ -static'

FROM scratch

COPY --from=builder /ws/ws /
ENTRYPOINT ["/ws"]
