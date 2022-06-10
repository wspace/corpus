FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/wspace/zrneely-rust whitespace
WORKDIR /home/whitespace
RUN cargo test
RUN cargo build --release
# builds: /home/whitespace/target/release/whitespace
