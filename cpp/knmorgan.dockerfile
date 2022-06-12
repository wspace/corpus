FROM alpine as builder

RUN apk add git make g++ gmp-dev
RUN git clone https://github.com/knmorgan/ws
WORKDIR /ws
RUN make CC='g++ -static'

FROM scratch as runner

COPY --from=builder /ws/ws /
ENTRYPOINT ["/ws"]
