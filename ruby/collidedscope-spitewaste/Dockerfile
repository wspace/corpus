FROM ruby:3-alpine

RUN apk add --no-cache make gcc musl-dev
WORKDIR /spitewaste
COPY spitewaste .
# Remove `git ls-files` call:
RUN sed -i'' "/s.files/ s|= .*|= Dir['**/*']|" spitewaste.gemspec
RUN bundle install
RUN rake test
ENTRYPOINT ["ruby", "-Ilib", "bin/spw"]
