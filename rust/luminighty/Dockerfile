FROM rust:1.76 AS builder

WORKDIR /rustws
COPY rustws .
RUN cargo test
