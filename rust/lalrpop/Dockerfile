FROM rust:1.76 AS builder

WORKDIR /lalrpop
COPY lalrpop .
RUN cargo test --package whitespace && \
    RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu --package whitespace

FROM scratch

COPY --from=builder /lalrpop/target/x86_64-unknown-linux-gnu/release/whitespace /
COPY lalrpop/doc/whitespace/hello.ws /
ENTRYPOINT ["/whitespace"]
