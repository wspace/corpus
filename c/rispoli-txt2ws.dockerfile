FROM alpine as builder

RUN apk add git make gcc musl-dev
RUN git clone https://github.com/rispoli/txt2ws
WORKDIR /txt2ws
RUN make

FROM scratch as runner

COPY --from=builder /txt2ws/txt2ws /
ENTRYPOINT ["/txt2ws"]
