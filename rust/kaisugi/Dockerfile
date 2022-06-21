FROM rust:1.61 AS builder

RUN git clone https://github.com/HelloRusk/WhitespaceInterpreter
WORKDIR /WhitespaceInterpreter
RUN rustc -C target-feature=+crt-static -o wi whitespace_interpreter.rs

FROM scratch

COPY --from=builder /WhitespaceInterpreter/wi /
ENTRYPOINT ["/wi"]
