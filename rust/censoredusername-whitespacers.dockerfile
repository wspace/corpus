FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/CensoredUsername/whitespace-rs
WORKDIR /home/whitespace-rs
RUN cargo build --release
RUN test -f /home/whitespace-rs/target/release/wsc
