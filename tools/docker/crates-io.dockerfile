FROM rust:1.61

# Cache crates.io index
WORKDIR /tmp
RUN echo 'package = { name = "tmp", version = "0.1.0" }' > Cargo.toml
RUN echo 'dependencies._ = { version = "*" }' >> Cargo.toml
RUN mkdir src
RUN echo 'fn main() {}' > src/main.rs
RUN cargo build | :
RUN rm -rf /tmp
