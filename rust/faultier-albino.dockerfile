FROM rust:1.61

RUN git clone https://github.com/faultier/albino /faultier-albino
WORKDIR /faultier-albino
RUN cargo build --release
RUN cargo build
