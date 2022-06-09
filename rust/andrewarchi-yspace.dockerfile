FROM wspace-corpus/crates-io

RUN git clone https://github.com/andrewarchi/yspace /andrewarchi-yspace
WORKDIR /andrewarchi-yspace
RUN cargo build --release
# builds: /andrewarchi-yspace/target/release/yspace
