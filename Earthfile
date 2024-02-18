VERSION 0.8

IMPORT ./c

build:
    FROM scratch
    BUILD c+build
    WORKDIR /corpus
    COPY c+build/corpus/c c
    SAVE ARTIFACT /corpus /corpus

docker:
    FROM ubuntu:24.04
    WORKDIR /corpus
    COPY +build/ /
    SAVE IMAGE wspace-corpus

docker-all:
    BUILD +docker
    BUILD c+docker-all
