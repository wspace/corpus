FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/Luminighty/rustws
WORKDIR /home/rustws
RUN cargo build --release
# builds: /home/rustws/target/release/librustws.rlib
