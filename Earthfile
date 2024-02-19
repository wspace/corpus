VERSION 0.8

IMPORT ./c
IMPORT ./clojure
IMPORT ./coq
IMPORT ./crystal
IMPORT ./csharp
IMPORT ./erlang
IMPORT ./go
IMPORT ./haskell
IMPORT ./idris
IMPORT ./java

build:
    FROM scratch
    BUILD c+build
    BUILD clojure+build
    BUILD coq+build
    BUILD crystal+build
    BUILD csharp+build
    BUILD erlang+build
    BUILD go+build
    BUILD haskell+build
    BUILD idris+build
    BUILD java+build
    WORKDIR /corpus
    COPY c+build/corpus/c c
    COPY clojure+build/corpus/clojure clojure
    COPY coq+build/corpus/coq coq
    COPY crystal+build/corpus/crystal crystal
    COPY csharp+build/corpus/csharp csharp
    COPY erlang+build/corpus/erlang erlang
    COPY go+build/corpus/go go
    COPY haskell+build/corpus/haskell haskell
    COPY idris+build/corpus/idris idris
    COPY java+build/corpus/java java
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
    BUILD crystal+docker-all
    BUILD csharp+docker-all
    BUILD erlang+docker-all
    BUILD go+docker-all
    BUILD haskell+docker-all
    BUILD idris+docker-all
    BUILD java+docker-all
