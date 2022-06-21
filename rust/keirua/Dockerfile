FROM wspace-corpus/rust AS builder

RUN git clone https://github.com/Keirua/whitespace-rs
WORKDIR /whitespace-rs
RUN cargo test
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /whitespace-rs/target/x86_64-unknown-linux-gnu/release/compiler /
COPY --from=builder /whitespace-rs/target/x86_64-unknown-linux-gnu/release/interpreter /
ENTRYPOINT ["/compiler"]
