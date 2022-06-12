FROM wspace-corpus/crates-io as builder

RUN git clone https://github.com/Keirua/whitespace-rs
WORKDIR /whitespace-rs
RUN cargo test
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /whitespace-rs/target/release/compiler /
COPY --from=builder /whitespace-rs/target/release/interpreter /
ENTRYPOINT ["/compiler"]
