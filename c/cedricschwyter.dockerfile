FROM alpine as builder

RUN apk add git cmake make gcc g++
RUN git clone https://github.com/D3PSI/whitespace-interpreter
WORKDIR /whitespace-interpreter
RUN cmake .
RUN make

FROM scratch as runner

COPY --from=builder /whitespace-interpreter/interpreter /
ENTRYPOINT ["/interpreter"]
