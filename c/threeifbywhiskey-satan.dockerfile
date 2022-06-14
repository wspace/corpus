FROM alpine AS builder

RUN apk add git make gcc musl-dev
RUN git clone https://github.com/threeifbywhiskey/satan
WORKDIR /satan
RUN make

FROM scratch

COPY --from=builder /satan/satan /
ENTRYPOINT ["/satan"]
