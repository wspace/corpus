FROM rust:1.61 as builder

RUN git clone https://github.com/HelloRusk/WhitespaceInterpreter
WORKDIR /WhitespaceInterpreter
RUN rustc -o wi whitespace_interpreter.rs

FROM scratch as runner

COPY --from=builder /WhitespaceInterpreter/wi /
ENTRYPOINT ["/wi"]
