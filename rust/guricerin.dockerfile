FROM wspace-corpus/rust AS builder

RUN git clone https://github.com/guricerin/esolangs
WORKDIR /esolangs/whitespace-rs
RUN cargo test
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /esolangs/whitespace-rs/target/x86_64-unknown-linux-gnu/release/whitespace-rs /
ENTRYPOINT ["/whitespace-rs"]
