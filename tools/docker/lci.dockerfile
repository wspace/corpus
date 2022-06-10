FROM alpine

RUN apk add git python3 cmake make gcc g++ readline-dev
RUN git clone https://github.com/justinmeza/lci --branch=future /lci
WORKDIR /lci
RUN python3 install.py
