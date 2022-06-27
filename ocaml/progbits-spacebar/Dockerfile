FROM ocaml/opam:debian AS builder

USER root
RUN apt-get update && \
    apt-get install -y --no-install-recommends bats && \
    rm -rf /var/lib/apt/lists/*

USER opam
RUN opam install dune menhir ounit2
COPY --chown=opam spacebar /home/opam/spacebar
WORKDIR /home/opam/spacebar
RUN eval $(opam env) && dune runtest
RUN eval $(opam env) && dune build

RUN sed -i'' 's,$(git rev-parse --show-toplevel),.,' test/functional.sh
# TODO: Needs wspace binary:
# RUN test/functional.sh

FROM debian

WORKDIR /spacebar
COPY --from=builder /home/opam/spacebar/_build/default/spacebar.exe .
COPY spacebar/test/*.c ./
ENTRYPOINT ["./spacebar.exe"]
