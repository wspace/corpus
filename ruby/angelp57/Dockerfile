FROM ruby:3-alpine

WORKDIR /whitespace
COPY whitespace .
RUN ruby whitespace.rb -Sn sample/hello.wss && \
    ruby whitespace.rb -Sn sample/quine.wss && \
    ruby whitespace.rb -Sn sample/quine406.wss && \
    ruby whitespace.rb -Sn sample/sum_cnum_seq.wss
ENTRYPOINT ["ruby", "whitespace.rb"]
# TODO: Handle Perl entrypoints
