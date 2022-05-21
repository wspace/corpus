FROM rust:1.61

RUN git clone https://github.com/guricerin/esolangs /guricerin-rust
WORKDIR /guricerin-rust
RUN cargo build --release --manifest-path=whitespace-rs/Cargo.toml
RUN cargo build --manifest-path=whitespace-rs/Cargo.toml
