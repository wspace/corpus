FROM haskell:8 AS builder

RUN apt-get update && \
    apt-get install -y libnuma-dev
RUN cabal update
RUN cabal install -j idris
WORKDIR /WS-idr
COPY WS-idr .
RUN idris --build wspace.ipkg

FROM scratch

COPY --from=builder /WS-idr/wspace /WS-idr/*.ws ./
ENTRYPOINT ["/wspace"]
