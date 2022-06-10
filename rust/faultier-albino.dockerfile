FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/faultier/albino
WORKDIR /home/albino
RUN cargo build --release
# builds: /home/albino/target/release/albino, /home/albino/target/release/albino-run, /home/albino/target/release/albino-build, /home/albino/target/release/albino-exec, /home/albino/target/release/albino-gen
