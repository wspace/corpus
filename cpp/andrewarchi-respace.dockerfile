FROM alpine as builder

RUN apk add git make g++
RUN git clone https://github.com/andrewarchi/respace
WORKDIR /respace
RUN make test
RUN make LDFLAGS=-static

FROM scratch as runner

COPY --from=builder /respace/respace /
ENTRYPOINT ["/respace"]
