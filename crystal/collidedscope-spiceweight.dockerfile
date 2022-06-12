FROM alpine as builder

RUN apk add git make crystal
RUN git clone https://github.com/collidedscope/spiceweight
WORKDIR /spiceweight
RUN make

FROM scratch as runner

COPY --from=builder /spiceweight/spwt /
ENTRYPOINT ["/spwt"]
