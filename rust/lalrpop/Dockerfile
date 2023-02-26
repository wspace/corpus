FROM rust:1.67 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
# Nightly for sparse cargo index
RUN rustup override set nightly
WORKDIR /lalrpop
COPY lalrpop .
RUN cargo test -p whitespace
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu -p whitespace

FROM scratch

COPY --from=builder /lalrpop/target/x86_64-unknown-linux-gnu/release/whitespace /
COPY lalrpop/doc/whitespace/hello.whitespace /
ENTRYPOINT ["/whitespace"]
