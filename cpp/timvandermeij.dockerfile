FROM alpine AS builder

RUN apk add git make g++
RUN git clone https://github.com/wspace/timvandermeij-cpp whitespace-interpreter
WORKDIR /whitespace-interpreter
RUN make

FROM scratch

COPY --from=builder /whitespace-interpreter/whitespace /
ENTRYPOINT ["/whitespace"]
