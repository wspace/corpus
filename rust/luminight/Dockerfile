FROM rust:1.68 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
WORKDIR /rustws
COPY rustws .
RUN cargo test
