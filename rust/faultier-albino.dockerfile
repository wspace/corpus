FROM wspace-corpus/crates-io as builder

RUN git clone https://github.com/faultier/albino
WORKDIR /albino
RUN cargo build --release

FROM scratch as runner

COPY --from=builder /albino/target/release/albino /
COPY --from=builder /albino/target/release/albino-run /
COPY --from=builder /albino/target/release/albino-build /
COPY --from=builder /albino/target/release/albino-exec /
COPY --from=builder /albino/target/release/albino-gen /
ENTRYPOINT ["/albino"]
