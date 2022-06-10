FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/Keirua/whitespace-rs
WORKDIR /home/whitespace-rs
RUN cargo build --release
# builds: /home/whitespace-rs/target/release/compiler, /home/whitespace-rs/target/release/interpreter
