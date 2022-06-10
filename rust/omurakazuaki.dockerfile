FROM rust:1.61

WORKDIR /home
RUN git clone https://github.com/wspace/omurakazuaki-rust whitespace
WORKDIR /home/whitespace
RUN cargo build --release
# builds: /home/whitespace/target/release/whitespace
