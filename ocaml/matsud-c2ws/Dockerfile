FROM ocaml/opam AS builder

USER opam
COPY --chown=opam c2ws /home/opam/c2ws
WORKDIR /home/opam/c2ws
RUN eval $(opam env) && make

FROM scratch

COPY --from=builder /home/opam/c2ws/c2ws /
ENTRYPOINT ["/c2ws"]
