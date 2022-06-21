FROM rust:1.61 AS builder

WORKDIR /WhitespaceVM
COPY WhitespaceVM .
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /WhitespaceVM/target/x86_64-unknown-linux-gnu/release/whitespace-vm /
ENTRYPOINT ["/whitespace-vm"]
