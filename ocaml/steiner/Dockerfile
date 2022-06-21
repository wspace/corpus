FROM ocaml/opam AS builder

USER opam
WORKDIR /home/opam
RUN git clone https://github.com/steiner26/Whitespace
WORKDIR /home/opam/Whitespace
RUN eval $(opam env) && make

FROM scratch

COPY --from=builder /home/opam/Whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
