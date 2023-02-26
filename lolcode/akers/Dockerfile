FROM alpine AS builder

RUN apk add --no-cache git cmake make gcc g++ readline-dev
RUN git clone https://github.com/justinmeza/lci --branch=future
WORKDIR /lci
RUN cmake .
RUN make
# RUN ctest

FROM alpine

RUN apk add --no-cache bash readline-dev
WORKDIR /lolcode
COPY --from=builder /lci/lci /usr/local/bin/
COPY lolcode/*.lol lolcode/*.ws lolcode/*.b run.bash ./
ENTRYPOINT ["./run.bash"]
