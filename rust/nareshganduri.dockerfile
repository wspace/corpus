FROM rust:1.61

WORKDIR /home
RUN git clone https://github.com/nareshganduri/WhitespaceVM
WORKDIR /home/WhitespaceVM
RUN cargo build --release
RUN test -f /home/WhitespaceVM/target/release/whitespace-vm
