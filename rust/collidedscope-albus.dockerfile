FROM rust:1.61

RUN git clone https://github.com/collidedscope/albus /collidedscope-albus
WORKDIR /collidedscope-albus
RUN cargo build --release
RUN cargo build
