FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/Keirua/whitespace-rs
WORKDIR /home/whitespace-rs
RUN cargo test
RUN cargo build --release
RUN test -f /home/whitespace-rs/target/release/compiler
RUN test -f /home/whitespace-rs/target/release/interpreter
