FROM rust:1.61

RUN git clone https://github.com/Jayshua/rust-whitespace /jayshua-rust
WORKDIR /jayshua-rust
RUN cargo build --release
# builds: /jayshua-rust/target/release/whitespace
