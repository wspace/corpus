FROM rust:1.68 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
WORKDIR /ws
COPY ws .
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /ws/target/x86_64-unknown-linux-gnu/release/ws /ws/examples/*.ws ./
ENTRYPOINT ["/ws"]
