FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev
WORKDIR /txt2ws
COPY txt2ws .
RUN make

FROM scratch

COPY --from=builder /txt2ws/txt2ws /
ENTRYPOINT ["/txt2ws"]
