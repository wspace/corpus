FROM ocaml/opam AS builder

USER opam
COPY --chown=opam Whitespace /home/opam/Whitespace
WORKDIR /home/opam/Whitespace
RUN eval $(opam env) && make

FROM scratch

COPY --from=builder /home/opam/Whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
