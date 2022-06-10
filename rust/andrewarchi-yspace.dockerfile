FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/andrewarchi/yspace
WORKDIR /home/yspace
RUN cargo build --release
# builds: /home/yspace/target/release/yspace
