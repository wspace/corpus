FROM alpine AS builder

RUN apk add --no-cache gcc musl-dev perl
WORKDIR /ws
COPY ws .
RUN perl parsegen.pl parse.def > gencode.c && \
    gcc -O3 -g -static -o wspace parse.c vm.c wspace.c

FROM scratch

COPY --from=builder /ws/wspace /
ENTRYPOINT ["/wspace"]
