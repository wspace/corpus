FROM rust:1.61 AS builder

RUN git clone https://github.com/wspace/omurakazuaki-rust whitespace
WORKDIR /whitespace
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /whitespace/target/x86_64-unknown-linux-gnu/release/whitespace /
ENTRYPOINT ["/whitespace"]
