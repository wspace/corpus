FROM rust:1.67 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
# Nightly for sparse cargo index
RUN rustup override set nightly
WORKDIR /albus
COPY albus .
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /albus/target/x86_64-unknown-linux-gnu/release/albus /
ENTRYPOINT ["/albus"]
