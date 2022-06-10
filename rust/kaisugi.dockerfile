FROM rust:1.61

WORKDIR /home
RUN git clone https://github.com/HelloRusk/WhitespaceInterpreter
WORKDIR /home/WhitespaceInterpreter
RUN rustc -o wi whitespace_interpreter.rs
# builds: /home/WhitespaceInterpreter/wi
