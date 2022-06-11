FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/guricerin/esolangs
WORKDIR /home/esolangs/whitespace-rs
RUN cargo test
RUN cargo build --release
RUN test -f /home/esolangs/whitespace-rs/target/release/whitespace-rs
