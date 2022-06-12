FROM wspace-corpus/crates-io as builder

RUN git clone https://github.com/wspace/zrneely-rust whitespace
WORKDIR /whitespace
RUN cargo test
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /whitespace/target/release/whitespace /
ENTRYPOINT ["/whitespace"]
