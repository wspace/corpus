FROM rust:1.61

RUN git clone https://github.com/Keirua/whitespace-rs /keirua-rust
WORKDIR /keirua-rust
RUN cargo build --release
# builds: /keirua-rust/target/release/compiler, /keirua-rust/target/release/interpreter
