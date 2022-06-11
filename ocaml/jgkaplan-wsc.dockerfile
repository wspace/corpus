FROM ocaml/opam

USER opam
RUN opam install ocamlbuild ocamlfind menhir
WORKDIR /home/opam
RUN git clone https://github.com/jgkaplan/whitespaceTranspiler
WORKDIR /home/opam/whitespaceTranspiler
RUN eval $(opam env) && make
RUN test -f /home/opam/whitespaceTranspiler/main.byte
