FROM alpine as builder

RUN apk add git cmake make gcc g++ readline-dev
RUN git clone https://github.com/wspace/akers-lolcode lolcode
RUN git clone https://github.com/justinmeza/lci --branch=future
WORKDIR /lci
RUN cmake .
RUN make
# RUN ctest

FROM alpine as runner

RUN apk add readline-dev
WORKDIR /lolcode
COPY --from=builder /lci/lci /usr/local/bin/
COPY --from=builder /lolcode/*.lol /lolcode/*.ws /lolcode/*.b ./
ENTRYPOINT ["/usr/local/bin/lci", "whitespace.lol"]
