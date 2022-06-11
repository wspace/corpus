FROM alpine

RUN apk add git make gcc musl-dev
WORKDIR /home
RUN git clone https://github.com/subgeniuskitty/vvhitespace
WORKDIR /home/vvhitespace
RUN test -f /home/vvhitespace/syntax_highlighting/vvhitespace.vim
RUN make
RUN test -f /home/vvhitespace/vvc
RUN test -f /home/vvhitespace/vvi
