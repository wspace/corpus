FROM alpine as builder

RUN apk add git make gcc musl-dev perl
RUN git clone https://github.com/wspace/hogelog-c ws
WORKDIR /ws
RUN perl parsegen.pl parse.def > gencode.c
RUN make

FROM scratch as runner

COPY --from=builder /ws/wspace /
ENTRYPOINT ["/wspace"]
