FROM alpine as builder

RUN apk add git make gcc musl-dev
RUN git clone https://github.com/subgeniuskitty/vvhitespace
WORKDIR /vvhitespace
RUN make

FROM scratch as runner

COPY --from=builder /vvhitespace/vvc /
COPY --from=builder /vvhitespace/vvi /
ENTRYPOINT ["/vvhitespace.vim"]
