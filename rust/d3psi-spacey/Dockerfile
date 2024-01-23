FROM rust:1.75 AS builder

WORKDIR /spacey
COPY spacey .
RUN rustup override set nightly-2023-06-28 && \
    cargo test && \
    RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /spacey/target/x86_64-unknown-linux-gnu/release/spacey /
COPY spacey/resources/ws/*.ws /
ENTRYPOINT ["/spacey"]
