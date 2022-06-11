FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/andrewarchi/yspace
WORKDIR /home/yspace
RUN cargo test
RUN cargo build --release
RUN test -f /home/yspace/target/release/yspace
