FROM alpine AS builder

RUN apk add --no-cache make g++ flex bison
WORKDIR /whitespace_compiler
COPY whitespace_compiler .
RUN make

FROM scratch

COPY --from=builder /whitespace_compiler/compile /
ENTRYPOINT ["/compile"]
