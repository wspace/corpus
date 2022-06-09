FROM rust:1.61

RUN git clone https://github.com/wspace/omurakazuaki-rust /omurakazuaki-rust
WORKDIR /omurakazuaki-rust
RUN cargo build --release
# builds: /omurakazuaki-rust/target/release/whitespace
