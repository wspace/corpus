FROM alpine AS builder

RUN apk add git cmake make gcc g++ readline-dev
RUN git clone https://github.com/justinmeza/lci --branch=future
WORKDIR /lci
RUN cmake .
RUN make
# RUN ctest

FROM alpine

RUN apk add readline-dev
WORKDIR /lolcode
COPY --from=builder /lci/lci /usr/local/bin/
COPY lolcode/*.lol lolcode/*.ws lolcode/*.b ./
ENTRYPOINT ["lci", "whitespace.lol"]
