FROM ocaml/opam

USER opam
RUN opam install dune core ounit2
WORKDIR /home/opam
RUN git clone https://github.com/progbits/blank
WORKDIR /home/opam/blank
RUN eval $(opam env) && dune build
RUN test -f /home/opam/blank/_build/default/blank.exe
