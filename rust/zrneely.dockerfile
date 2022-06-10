FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://gitlab.com/zrneely/whitespace
WORKDIR /home/whitespace
RUN cargo build --release
# builds: /home/whitespace/target/release/whitespace
