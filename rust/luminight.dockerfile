FROM wspace-corpus/rust AS builder

RUN git clone https://github.com/Luminighty/rustws
WORKDIR /rustws
RUN cargo test
