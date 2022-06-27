FROM ocaml/opam AS builder

USER opam
RUN opam install dune core ounit2
COPY --chown=opam blank /home/opam/blank
WORKDIR /home/opam/blank
RUN eval $(opam env) && dune runtest
RUN eval $(opam env) && dune build

FROM scratch

COPY --from=builder /home/opam/blank/_build/default/blank.exe /
ENTRYPOINT ["/blank.exe"]
