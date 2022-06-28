FROM ocaml/opam:debian

USER opam
COPY --chown=opam Whitespace /home/opam/Whitespace
WORKDIR /home/opam/Whitespace
RUN eval $(opam env) && make
ENTRYPOINT ["./whitespace"]
