FROM ruby:3-alpine

WORKDIR /whitespace-ruby
COPY whitespace-ruby .
# Fixes: `define_method': tried to create Proc object without a block
RUN sed -i 's/define_method ops\[1\]/define_method(ops[1])/' wsv.rb
# sample/facws.rb or sample/fibws.rb
ENTRYPOINT ["ruby"]
