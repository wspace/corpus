FROM wspace-corpus/crates-io

RUN git clone https://github.com/guricerin/esolangs /guricerin-rust
WORKDIR /guricerin-rust/whitespace-rs
RUN cargo build --release
# builds: /guricerin-rust/whitespace-rs/target/release/whitespace-rs
