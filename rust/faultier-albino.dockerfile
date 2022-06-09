FROM wspace-corpus/crates-io

RUN git clone https://github.com/faultier/albino /faultier-albino
WORKDIR /faultier-albino
RUN cargo build --release
# builds: /faultier-albino/target/release/albino, /faultier-albino/target/release/albino-run, /faultier-albino/target/release/albino-build, /faultier-albino/target/release/albino-exec, /faultier-albino/target/release/albino-gen
