FROM alpine AS builder

RUN apk add make gcc musl-dev
WORKDIR /satan
COPY satan .
RUN make

FROM scratch

COPY --from=builder /satan/satan /
ENTRYPOINT ["/satan"]
