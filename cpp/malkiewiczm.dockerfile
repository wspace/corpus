FROM alpine AS builder

RUN apk add git make g++ flex bison
RUN git clone https://github.com/malkiewiczm/whitespace_compiler
WORKDIR /whitespace_compiler
RUN make

FROM scratch

COPY --from=builder /whitespace_compiler/compile /
ENTRYPOINT ["/compile"]
