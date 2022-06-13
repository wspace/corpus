FROM rust:1.61 as builder

RUN git clone https://github.com/nareshganduri/WhitespaceVM
WORKDIR /WhitespaceVM
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch as runner

COPY --from=builder /WhitespaceVM/target/x86_64-unknown-linux-gnu/release/whitespace-vm /
ENTRYPOINT ["/whitespace-vm"]
