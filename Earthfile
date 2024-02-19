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
    SAVE ARTIFACT /corpus /corpus

docker:
    FROM ubuntu:24.04
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
