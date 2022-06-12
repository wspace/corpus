FROM rust:1.61 as builder

RUN git clone https://github.com/Jayshua/rust-whitespace
WORKDIR /rust-whitespace
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /rust-whitespace/target/release/whitespace /
ENTRYPOINT ["/whitespace"]
