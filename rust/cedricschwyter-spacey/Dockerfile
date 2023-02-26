FROM rust:1.67 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
RUN rustup update nightly
WORKDIR /spacey
COPY spacey .
RUN cargo test
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /spacey/target/x86_64-unknown-linux-gnu/release/spacey /
COPY spacey/resources/ws/*.ws /
ENTRYPOINT ["/spacey"]
