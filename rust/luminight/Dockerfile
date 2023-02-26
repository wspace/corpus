FROM rust:1.67 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
# Nightly for sparse cargo index
RUN rustup override set nightly
WORKDIR /rustws
COPY rustws .
RUN cargo test
