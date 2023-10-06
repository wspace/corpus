FROM rust:1.73 AS builder

RUN rustup update nightly
WORKDIR /spacey
COPY spacey .
# Build now fails from unstable features
RUN cargo test
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /spacey/target/x86_64-unknown-linux-gnu/release/spacey /
COPY spacey/resources/ws/*.ws /
ENTRYPOINT ["/spacey"]
