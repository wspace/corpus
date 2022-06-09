FROM wspace-corpus/crates-io

RUN git clone https://github.com/collidedscope/albus /collidedscope-albus
WORKDIR /collidedscope-albus
RUN cargo build --release
# builds: /collidedscope-albus/target/release/albus
