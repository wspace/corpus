FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev perl
WORKDIR /ws
COPY ws .
RUN perl parsegen.pl parse.def > gencode.c
RUN make

FROM scratch

COPY --from=builder /ws/wspace /
ENTRYPOINT ["/wspace"]
