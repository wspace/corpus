FROM wspace-corpus/rust AS builder

WORKDIR /spacey
COPY spacey .
RUN cargo test
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /spacey/target/x86_64-unknown-linux-gnu/release/spacey /
COPY --from=builder /spacey/ws/*.ws ./
ENTRYPOINT ["/spacey"]
