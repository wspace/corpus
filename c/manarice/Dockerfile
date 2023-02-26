FROM alpine AS builder

RUN apk add --no-cache make gcc musl-dev
WORKDIR /utils/limelib
COPY limelib .
RUN make
WORKDIR /utils/whitespace
COPY whitespace .
RUN make

FROM scratch

COPY --from=builder /utils/whitespace/lwsvm /
COPY --from=builder /utils/whitespace/lwsa /
ENTRYPOINT ["/lwsvm"]
