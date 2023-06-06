FROM rust:1.70 AS builder

WORKDIR /rustws
COPY rustws .
RUN cargo test
