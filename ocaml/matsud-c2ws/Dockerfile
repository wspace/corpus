FROM ocaml/opam:debian AS builder

USER opam
COPY --chown=opam c2ws /home/opam/c2ws
WORKDIR /home/opam/c2ws
RUN eval $(opam env) && \
    ocamllex -q mylexer.mll && \
    ocamlyacc myparser.mly && \
    ocamlopt -O3 -o c2ws $(ocamldep -native -sort *.ml *.mli)

FROM debian

WORKDIR /c2ws
COPY --from=builder /home/opam/c2ws/c2ws .
ENTRYPOINT ["./c2ws"]
