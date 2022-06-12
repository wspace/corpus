FROM wspace-corpus/crates-io as builder

RUN git clone https://github.com/CensoredUsername/whitespace-rs
WORKDIR /whitespace-rs
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /whitespace-rs/target/release/wsc /
ENTRYPOINT ["/wsc"]
