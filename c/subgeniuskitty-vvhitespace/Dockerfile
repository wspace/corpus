FROM alpine AS builder

RUN apk add make gcc musl-dev
WORKDIR /vvhitespace
COPY vvhitespace .
RUN make

FROM scratch

COPY --from=builder /vvhitespace/vvc /
COPY --from=builder /vvhitespace/vvi /
ENTRYPOINT ["/vvhitespace.vim"]
