FROM wspace-corpus/crates-io as builder

RUN git clone https://github.com/collidedscope/albus
WORKDIR /albus
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /albus/target/release/albus /
ENTRYPOINT ["/albus"]
