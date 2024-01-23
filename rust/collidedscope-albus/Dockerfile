FROM rust:1.75 AS builder

WORKDIR /albus
COPY albus .
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /albus/target/x86_64-unknown-linux-gnu/release/albus /
ENTRYPOINT ["/albus"]
