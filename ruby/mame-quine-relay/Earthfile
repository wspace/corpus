FROM ruby:3-alpine

WORKDIR /quine-relay
COPY quine-relay/vendor/whitespace.rb .
ENTRYPOINT ["ruby", "whitespace.rb"]
