FROM wspace-corpus/rust as builder

RUN git clone https://github.com/andrewarchi/yspace
WORKDIR /yspace
RUN cargo test
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /yspace/target/release/yspace /
ENTRYPOINT ["/yspace"]
