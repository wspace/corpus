FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/guricerin/esolangs
WORKDIR /home/esolangs/whitespace-rs
RUN cargo build --release
# builds: /home/esolangs/whitespace-rs/target/release/whitespace-rs
