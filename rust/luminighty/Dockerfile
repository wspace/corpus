FROM rust:1.73 AS builder

WORKDIR /rustws
COPY rustws .
RUN cargo test
