FROM ocaml/opam:debian AS builder

USER opam
RUN opam install ocamlbuild ocamlfind menhir
COPY --chown=opam whitespaceTranspiler /home/opam/whitespaceTranspiler
WORKDIR /home/opam/whitespaceTranspiler
RUN eval $(opam env) && \
    ocamlbuild -use-ocamlfind -use-menhir -I src main.native

FROM debian

WORKDIR /whitespaceTranspiler
COPY --from=builder /home/opam/whitespaceTranspiler/_build/src/main.native .
COPY whitespaceTranspiler/lib lib
COPY whitespaceTranspiler/test test
ENTRYPOINT ["./main.native"]
