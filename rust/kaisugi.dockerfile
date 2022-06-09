FROM rust:1.61

RUN git clone https://github.com/HelloRusk/WhitespaceInterpreter /kaisugi-rust
WORKDIR /kaisugi-rust
RUN rustc -o wi whitespace_interpreter.rs
# builds: /kaisugi-rust/wi
