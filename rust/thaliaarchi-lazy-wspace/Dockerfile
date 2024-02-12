FROM coqorg/coq AS theories

USER coq
COPY --chown=coq lazy-wspace/theories /home/coq/lazy-wspace/theories
COPY --chown=coq lazy-wspace/_CoqProject lazy-wspace/Makefile /home/coq/lazy-wspace
WORKDIR /home/coq/lazy-wspace
RUN make
RUN touch build_stamp

FROM rust:1.76 AS builder

WORKDIR /lazy-wspace
COPY lazy-wspace .
# Exclude package wspace-syntax-test, because it requires GHC for comparing
# against Haskell behavior. TODO: Enable it.
RUN cargo test --package lazy-wspace \
               --package lazy-wspace-macros \
               --package wspace-syntax && \
    RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu
# Force a dependency on the proofs, since no artifacts are used.
COPY --from=theories /home/coq/lazy-wspace/build_stamp .

FROM scratch

COPY --from=builder /lazy-wspace/target/x86_64-unknown-linux-gnu/release/compile \
                    /lazy-wspace/target/x86_64-unknown-linux-gnu/release/wspace \
                    /lazy-wspace/target/x86_64-unknown-linux-gnu/release/wspace-syntax /
COPY lazy-wspace/tests /tests
ENTRYPOINT ["/compile"]
