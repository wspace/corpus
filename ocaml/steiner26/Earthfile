FROM ocaml/opam:debian AS builder

USER opam
COPY --chown=opam Whitespace /home/opam/Whitespace
WORKDIR /home/opam/Whitespace
RUN eval $(opam env) && \
	  ocamllex -q lexer.mll && \
	  ocamlyacc parser.mly && \
    ocamlopt -O3 -o whitespace $(ocamldep -native -sort *.ml *.mli)

FROM debian

WORKDIR /Whitespace
COPY --from=builder /home/opam/Whitespace/whitespace .
COPY Whitespace/programs .
ENTRYPOINT ["./whitespace"]
