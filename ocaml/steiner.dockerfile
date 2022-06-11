FROM ocaml/opam

USER opam
WORKDIR /home/opam
RUN git clone https://github.com/steiner26/Whitespace
WORKDIR /home/opam/Whitespace
RUN eval $(opam env) && make
RUN test -f /home/opam/Whitespace/whitespace
