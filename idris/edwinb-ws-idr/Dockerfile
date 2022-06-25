FROM haskell:8 AS builder

RUN cabal update
RUN cabal install idris
WORKDIR /WS-idr
COPY WS-idr .
RUN idris --build wspace.ipkg

FROM scratch

COPY --from=builder /WS-idr/wspace /WS-idr/*.ws ./
ENTRYPOINT ["/wspace"]
