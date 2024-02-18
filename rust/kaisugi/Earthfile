FROM rust:1.76 AS builder

WORKDIR /WhitespaceInterpreter
COPY WhitespaceInterpreter .
RUN rustc -C target-feature=+crt-static -o wi whitespace_interpreter.rs

FROM scratch

COPY --from=builder /WhitespaceInterpreter/wi /
ENTRYPOINT ["/wi"]
