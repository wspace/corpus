FROM rust:1.61

WORKDIR /home
RUN git clone https://github.com/wspace/omurakazuaki-rust
WORKDIR /home/omurakazuaki-rust
RUN cargo build --release
# builds: /home/omurakazuaki-rust/target/release/whitespace
