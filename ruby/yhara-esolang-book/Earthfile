FROM ruby:3-alpine

WORKDIR /esolang-book-sources
COPY esolang-book-sources .
RUN rake test
ENTRYPOINT ["ruby", "whitespace/bin/whitespace"]
