FROM alpine AS builder

RUN apk add make g++
WORKDIR /wspace
COPY wspace .
RUN make

FROM scratch

COPY --from=builder /wspace/wspace /
COPY --from=builder /wspace/assembler /
ENTRYPOINT ["/wspace"]
