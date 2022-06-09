FROM rust:1.61

RUN git clone https://gitlab.com/zrneely/whitespace /zrneely-rust
WORKDIR /zrneely-rust
RUN cargo build --release
# builds: /zrneely-rust/target/release/whitespace
