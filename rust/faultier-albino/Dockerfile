FROM rust:1.75 AS builder

WORKDIR /albino
COPY albino .
RUN RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu

FROM scratch

COPY --from=builder /albino/target/x86_64-unknown-linux-gnu/release/albino /
COPY --from=builder /albino/target/x86_64-unknown-linux-gnu/release/albino-run /
COPY --from=builder /albino/target/x86_64-unknown-linux-gnu/release/albino-build /
COPY --from=builder /albino/target/x86_64-unknown-linux-gnu/release/albino-exec /
COPY --from=builder /albino/target/x86_64-unknown-linux-gnu/release/albino-gen /
ENTRYPOINT ["/albino"]
