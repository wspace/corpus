FROM rust:1.61

RUN git clone https://github.com/CensoredUsername/whitespace-rs /censoredusername-whitespacers
WORKDIR /censoredusername-whitespacers
RUN cargo build --release
RUN cargo build
