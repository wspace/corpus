FROM rust:1.68 AS builder

ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL sparse
WORKDIR /WhitespaceInterpreter
COPY WhitespaceInterpreter .
RUN rustc -C target-feature=+crt-static -o wi whitespace_interpreter.rs

FROM scratch

COPY --from=builder /WhitespaceInterpreter/wi /
ENTRYPOINT ["/wi"]
