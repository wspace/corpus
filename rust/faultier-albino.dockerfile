FROM wspace-corpus/crates-io

WORKDIR /home
RUN git clone https://github.com/faultier/albino
WORKDIR /home/albino
RUN cargo build --release
RUN test -f /home/albino/target/release/albino
RUN test -f /home/albino/target/release/albino-run
RUN test -f /home/albino/target/release/albino-build
RUN test -f /home/albino/target/release/albino-exec
RUN test -f /home/albino/target/release/albino-gen
