FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev
WORKDIR /blacktime
COPY blacktime .
RUN make

FROM scratch

COPY --from=builder /blacktime/blacktime /
ENTRYPOINT ["/blacktime"]
