FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev flex bison
WORKDIR /ws-compiler
COPY ws-compiler/flex-bison/262_267_256 .
RUN make

FROM scratch

COPY --from=builder /ws-compiler/ws /
ENTRYPOINT ["/ws"]
