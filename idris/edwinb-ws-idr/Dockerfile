FROM haskell:8-slim-buster AS builder

RUN cabal update
RUN cabal install -j idris
WORKDIR /WS-idr
COPY WS-idr .
RUN idris --build wspace.ipkg

FROM debian:buster-slim

WORKDIR /WS-idr
COPY --from=builder /WS-idr/wspace /WS-idr/*.ws ./
ENTRYPOINT ["./wspace"]
