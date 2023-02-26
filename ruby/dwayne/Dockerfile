FROM ruby:3-alpine

RUN apk add --no-cache make gcc musl-dev
WORKDIR /whitespace-ruby
COPY whitespace-ruby .
# Remove `git ls-files` calls and loosen dependency versions:
RUN sed -i'' \
        -e "/s\.executables/ s,= .*,= ['whitespace']," \
        -e "/s\.files/       s,= .*,= Dir['**/*']," \
        -e "/s\.test_files/  s,= .*,= Dir['test/**/*.rb']," \
        -e 's/~>/>=/' \
        whitespace-ruby.gemspec
RUN bundle install
RUN rake test
ENTRYPOINT ["ruby", "bin/whitespace"]
