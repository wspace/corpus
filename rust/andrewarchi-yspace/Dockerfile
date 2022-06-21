FROM wspace-corpus/rust AS builder

RUN git clone https://github.com/andrewarchi/yspace
WORKDIR /yspace
RUN cargo test
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /yspace/target/x86_64-unknown-linux-gnu/release/yspace /
ENTRYPOINT ["/yspace"]
