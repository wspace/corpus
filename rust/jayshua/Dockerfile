FROM rust:1.70 AS builder

WORKDIR /rust-whitespace
COPY rust-whitespace .
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /rust-whitespace/target/x86_64-unknown-linux-gnu/release/whitespace /
ENTRYPOINT ["/whitespace"]
