FROM alpine AS builder

RUN apk add cmake make gcc g++
WORKDIR /whitespace-interpreter
COPY whitespace-interpreter .
RUN cmake .
RUN make

FROM scratch

COPY --from=builder /whitespace-interpreter/interpreter /
ENTRYPOINT ["/interpreter"]
