FROM wspace-corpus/rust as builder

RUN git clone https://github.com/Luminighty/rustws
WORKDIR /rustws
RUN cargo test
