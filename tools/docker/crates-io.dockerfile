FROM rust:1.61

# Cache crates.io index
WORKDIR /stub
RUN echo 'package = { name = "stub", version = "0.1.0" }' > Cargo.toml
# Depend on a non-existent crate so that Cargo will update the crates.io index,
# but not fetch any crates.
RUN echo 'dependencies._ = { version = "*" }' >> Cargo.toml
RUN mkdir src
RUN echo 'fn main() {}' > src/main.rs
RUN cargo build | :
RUN rm -rf /stub
WORKDIR /
