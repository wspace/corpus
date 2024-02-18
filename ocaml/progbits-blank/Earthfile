FROM ocaml/opam:debian AS builder

USER opam
RUN opam install dune core ounit2
COPY --chown=opam blank /home/opam/blank
WORKDIR /home/opam/blank
RUN eval $(opam env) && \
    dune runtest && \
    dune build --profile release

FROM debian

WORKDIR /blank
COPY --from=builder /home/opam/blank/_build/default/blank.exe ./
COPY blank/examples .
ENTRYPOINT ["./blank.exe"]
