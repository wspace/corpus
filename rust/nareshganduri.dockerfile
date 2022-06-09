FROM rust:1.61

RUN git clone https://github.com/nareshganduri/WhitespaceVM /nareshganduri-rust
WORKDIR /nareshganduri-rust
RUN cargo build --release
# builds: /nareshganduri-rust/target/release/whitespace-vm
