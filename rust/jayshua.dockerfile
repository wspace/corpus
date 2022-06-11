FROM rust:1.61

WORKDIR /home
RUN git clone https://github.com/Jayshua/rust-whitespace
WORKDIR /home/rust-whitespace
RUN cargo build --release
RUN test -f /home/rust-whitespace/target/release/whitespace
