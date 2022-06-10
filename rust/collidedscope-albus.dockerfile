FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/collidedscope/albus
WORKDIR /home/albus
RUN cargo build --release
# builds: /home/albus/target/release/albus
