FROM rust:1.68 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
WORKDIR /whitespace
COPY whitespace .
RUN cargo test
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /whitespace/target/x86_64-unknown-linux-gnu/release/whitespace /
ENTRYPOINT ["/whitespace"]
