VERSION 0.8

IMPORT ./c
IMPORT ./clojure
IMPORT ./coq
IMPORT ./cpp
IMPORT ./crystal
IMPORT ./csharp
IMPORT ./erlang
IMPORT ./go
IMPORT ./haskell
IMPORT ./idris
IMPORT ./java
IMPORT ./javascript
IMPORT ./jq
IMPORT ./kotlin
IMPORT ./lolcode
IMPORT ./lua
IMPORT ./ocaml
IMPORT ./rust
IMPORT ./typescript

build:
    FROM scratch
    BUILD c+build
    BUILD clojure+build
    BUILD coq+build
    BUILD cpp+build
    BUILD crystal+build
    BUILD csharp+build
    BUILD erlang+build
    BUILD go+build
    BUILD haskell+build
    BUILD idris+build
    BUILD java+build
    BUILD javascript+build
    BUILD jq+build
    BUILD kotlin+build
    BUILD lolcode+build
    BUILD lua+build
    BUILD ocaml+build
    BUILD rust+build
    BUILD typescript+build
    WORKDIR /corpus
    COPY c+build/corpus/c c
    COPY clojure+build/corpus/clojure clojure
    COPY coq+build/corpus/coq coq
    COPY cpp+build/corpus/cpp cpp
    COPY crystal+build/corpus/crystal crystal
    COPY csharp+build/corpus/csharp csharp
    COPY erlang+build/corpus/erlang erlang
    COPY go+build/corpus/go go
    COPY haskell+build/corpus/haskell haskell
    COPY idris+build/corpus/idris idris
    COPY java+build/corpus/java java
    COPY javascript+build/corpus/javascript javascript
    COPY jq+build/corpus/jq jq
    COPY kotlin+build/corpus/kotlin kotlin
    COPY lolcode+build/corpus/lolcode lolcode
    COPY lua+build/corpus/lua lua
    COPY ocaml+build/corpus/ocaml ocaml
    COPY rust+build/corpus/rust rust
    COPY typescript+build/corpus/rust rust
    SAVE ARTIFACT /corpus /corpus

# Dependencies:
#   c: -
#   clojure: jre-21
#   coq: -
#   cpp: -
#     cpp/buyoh-nospace: ruby
#   crystal: -
#   csharp: mono-runtime
#     csharp/nicolasff-spacesharp: mono-runtime libmono-compilerservices-symbolwriter4.0-cil
#   erlang: erlang
#   go: -
#   haskell: TODO
#   idris: ?
#   java: jre-21
#   javascript: node
#   jq: jq bash
#   kotlin: jre-21
#   lolcode: bash readline-dev
#   lua: lua5.4
#   ocaml: -
#   rust: -
#   typescript: TODO
docker:
    FROM ubuntu:24.04
    RUN apt-get update && \
        DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
            mono-runtime libmono-compilerservices-symbolwriter4.0-cil \
        && \
        rm -rf /var/lib/apt/lists/*
    WORKDIR /corpus
    COPY +build/ /
    SAVE IMAGE wspace-corpus

docker-all:
    BUILD +docker
    BUILD c+docker-all
    BUILD clojure+docker-all
    BUILD coq+docker-all
    BUILD cpp+docker-all
    BUILD crystal+docker-all
    BUILD csharp+docker-all
    BUILD erlang+docker-all
    BUILD go+docker-all
    BUILD haskell+docker-all
    BUILD idris+docker-all
    BUILD java+docker-all
    BUILD javascript+docker-all
    BUILD jq+docker-all
    BUILD kotlin+docker-all
    BUILD lolcode+docker-all
    BUILD lua+docker-all
    BUILD ocaml+docker-all
    BUILD rust+docker-all
    BUILD typescript+docker-all
