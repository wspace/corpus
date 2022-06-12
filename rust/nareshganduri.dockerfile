FROM rust:1.61 as builder

RUN git clone https://github.com/nareshganduri/WhitespaceVM
WORKDIR /WhitespaceVM
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /WhitespaceVM/target/release/whitespace-vm /
ENTRYPOINT ["/whitespace-vm"]
