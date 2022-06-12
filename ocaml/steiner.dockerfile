FROM ocaml/opam as builder

USER opam
WORKDIR /home/opam
RUN git clone https://github.com/steiner26/Whitespace
WORKDIR /home/opam/Whitespace
RUN eval $(opam env) && make

FROM scratch as runner

COPY --from=builder /home/opam/Whitespace/whitespace /
ENTRYPOINT ["/whitespace"]
