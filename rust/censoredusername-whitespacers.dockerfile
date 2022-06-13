FROM wspace-corpus/rust as builder

RUN git clone https://github.com/CensoredUsername/whitespace-rs
WORKDIR /whitespace-rs
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch as runner

COPY --from=builder /whitespace-rs/target/x86_64-unknown-linux-gnu/release/wsc /
ENTRYPOINT ["/wsc"]
