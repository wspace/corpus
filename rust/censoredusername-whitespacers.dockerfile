FROM wspace-corpus/crates-io

RUN git clone https://github.com/CensoredUsername/whitespace-rs /censoredusername-whitespacers
WORKDIR /censoredusername-whitespacers
RUN cargo build --release
# builds: /censoredusername-whitespacers/target/release/wsc
