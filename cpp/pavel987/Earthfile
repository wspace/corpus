FROM alpine AS builder

RUN apk add --no-cache make g++
WORKDIR /wspace
COPY wspace .
RUN make

FROM scratch

COPY --from=builder /wspace/wspace /
COPY --from=builder /wspace/assembler /
ENTRYPOINT ["/wspace"]
