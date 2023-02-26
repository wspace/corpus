FROM rust:1.67 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
# Nightly for sparse cargo index
RUN rustup override set nightly
WORKDIR /whitespace-rs
COPY whitespace-rs .
RUN cargo test
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /whitespace-rs/target/x86_64-unknown-linux-gnu/release/compiler /
COPY --from=builder /whitespace-rs/target/x86_64-unknown-linux-gnu/release/interpreter /
ENTRYPOINT ["/compiler"]
