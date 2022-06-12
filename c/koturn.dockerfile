FROM alpine as builder

RUN apk add git make gcc musl-dev
RUN git clone https://github.com/wspace/koturn-c Whitespace
WORKDIR /Whitespace
RUN make

FROM scratch as runner

COPY --from=builder /Whitespace/whitespace.out /
ENTRYPOINT ["/whitespace.out"]
