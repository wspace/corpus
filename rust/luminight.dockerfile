FROM wspace-corpus/crates-io

RUN git clone https://github.com/Luminighty/rustws /luminight-rust
WORKDIR /luminight-rust
RUN cargo build --release
# builds: /luminight-rust/target/release/librustws.rlib
