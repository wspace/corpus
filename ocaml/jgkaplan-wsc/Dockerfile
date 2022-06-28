FROM ocaml/opam:debian

USER opam
RUN opam install ocamlbuild ocamlfind menhir
COPY --chown=opam whitespaceTranspiler /home/opam/whitespaceTranspiler
WORKDIR /home/opam/whitespaceTranspiler
RUN eval $(opam env) && make
ENTRYPOINT ["./main.byte"]
