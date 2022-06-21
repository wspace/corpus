FROM alpine AS builder

RUN apk add git make gcc musl-dev
RUN git clone https://github.com/threeifbywhiskey/blacktime
WORKDIR /blacktime
RUN make

FROM scratch

COPY --from=builder /blacktime/blacktime /
ENTRYPOINT ["/blacktime"]
