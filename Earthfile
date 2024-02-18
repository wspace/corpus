VERSION 0.8

IMPORT ./c
IMPORT ./clojure
IMPORT ./coq

build:
    FROM scratch
    BUILD c+build
    BUILD clojure+build
    BUILD coq+build
    WORKDIR /corpus
    COPY c+build/corpus/c c
    COPY clojure+build/corpus/clojure clojure
    COPY coq+build/corpus/coq coq
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
