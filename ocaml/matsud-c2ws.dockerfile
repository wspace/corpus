FROM ocaml/opam AS builder

USER opam
WORKDIR /home/opam
RUN git clone https://github.com/matsud224/c2ws
WORKDIR /home/opam/c2ws
RUN eval $(opam env) && make

FROM scratch

COPY --from=builder /home/opam/c2ws/c2ws /
ENTRYPOINT ["/c2ws"]
