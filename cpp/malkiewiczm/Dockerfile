FROM ubuntu AS builder

WORKDIR /whitespace_compiler
COPY whitespace_compiler .
RUN <<EOF
set -e

apt-get update
apt-get install -y --no-install-recommends make g++ flex bison
rm -rf /var/lib/apt/lists/*

make

./compile examples/spoj.txt > examples/spoj.ws
./compile -a examples/spoj.txt > examples/spoj_a.ws
./compile -i examples/spoj.txt > examples/spoj_i.ws
./compile -a -i examples/spoj.txt > examples/spoj_ai.ws

./compile examples/spoj.txt > examples/spoj.ws
./compile -a examples/spoj.txt > examples/spoj_a.ws
./compile -i examples/spoj.txt > examples/spoj_i.ws
./compile -a -i examples/spoj.txt > examples/spoj_ai.ws

./compile examples/spoj2.txt > examples/spoj2.ws
./compile -a examples/spoj2.txt > examples/spoj2_a.ws
./compile -i examples/spoj2.txt > examples/spoj2_i.ws
./compile -a -i examples/spoj2.txt > examples/spoj2_ai.ws

./compile examples/test.txt > examples/test.ws
./compile -a examples/test.txt > examples/test_a.ws
./compile -i examples/test.txt > examples/test_i.ws
./compile -a -i examples/test.txt > examples/test_ai.ws

EOF

FROM ubuntu

RUN apt-get update && \
    apt-get install -y --no-install-recommends flex bison && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /whitespace_compiler/compile /
COPY --from=builder /whitespace_compiler/examples /examples
ENTRYPOINT ["/compile"]
