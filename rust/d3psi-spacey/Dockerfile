FROM rust:1.76 AS builder

WORKDIR /spacey
COPY spacey .
RUN <<EOF
set -e

rustup override set nightly-2023-06-28
# Pin the version of clap, because 4.5.0 updated the MSRV from 1.70 to 1.74 and
# the nightly version is 1.72.
cargo update
cargo update -p clap@4.5.0 --precise 4.4.18

cargo test
RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

EOF

FROM scratch

COPY --from=builder /spacey/target/x86_64-unknown-linux-gnu/release/spacey /
COPY spacey/resources/ws/*.ws /
ENTRYPOINT ["/spacey"]
