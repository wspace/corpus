FROM wspace-corpus/rust as builder

RUN git clone https://github.com/guricerin/esolangs
WORKDIR /esolangs/whitespace-rs
RUN cargo test
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /esolangs/whitespace-rs/target/release/whitespace-rs /
ENTRYPOINT ["/whitespace-rs"]
