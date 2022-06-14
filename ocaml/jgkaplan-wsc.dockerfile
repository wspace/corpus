FROM ocaml/opam AS builder

USER opam
RUN opam install ocamlbuild ocamlfind menhir
WORKDIR /home/opam
RUN git clone https://github.com/jgkaplan/whitespaceTranspiler
WORKDIR /home/opam/whitespaceTranspiler
RUN eval $(opam env) && make

FROM scratch

COPY --from=builder /home/opam/whitespaceTranspiler/main.byte /
ENTRYPOINT ["/main.byte"]
