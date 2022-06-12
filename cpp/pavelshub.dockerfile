FROM alpine as builder

RUN apk add git make g++
RUN git clone https://github.com/wspace/pavelshub-cpp wspace
WORKDIR /wspace
RUN make

FROM scratch as runner

COPY --from=builder /wspace/wspace /
COPY --from=builder /wspace/assembler /
ENTRYPOINT ["/wspace"]
