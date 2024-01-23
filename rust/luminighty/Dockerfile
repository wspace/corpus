FROM rust:1.75 AS builder

WORKDIR /rustws
COPY rustws .
RUN cargo test
