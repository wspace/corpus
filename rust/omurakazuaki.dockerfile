FROM rust:1.61 as builder

RUN git clone https://github.com/wspace/omurakazuaki-rust whitespace
WORKDIR /whitespace
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /whitespace/target/release/whitespace /
ENTRYPOINT ["/whitespace"]
