FROM ocaml/opam

USER opam
WORKDIR /home/opam
RUN git clone https://github.com/matsud224/c2ws
WORKDIR /home/opam/c2ws
RUN eval $(opam env) && make
RUN test -f /home/opam/c2ws/c2ws
