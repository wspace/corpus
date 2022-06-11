FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/collidedscope/albus
WORKDIR /home/albus
RUN cargo build --release
RUN test -f /home/albus/target/release/albus
