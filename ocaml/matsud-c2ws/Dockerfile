FROM ocaml/opam:debian

USER opam
COPY --chown=opam c2ws /home/opam/c2ws
WORKDIR /home/opam/c2ws
RUN eval $(opam env) && make
ENTRYPOINT ["./c2ws"]
