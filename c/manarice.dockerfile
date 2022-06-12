FROM alpine as builder

RUN apk add git make gcc musl-dev
WORKDIR /utils
RUN git clone https://github.com/ManaRice/limelib
RUN git clone https://github.com/ManaRice/whitespace
WORKDIR /utils/limelib
RUN make
WORKDIR /utils/whitespace
RUN make

FROM scratch as runner

COPY --from=builder /utils/whitespace/lwsvm /
COPY --from=builder /utils/whitespace/lwsa /
ENTRYPOINT ["/lwsvm"]
