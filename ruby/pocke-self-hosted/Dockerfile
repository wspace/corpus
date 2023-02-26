FROM wspace-corpus/go/pocke-gows AS gows

FROM wspace-corpus/ruby/pocke-akaza

RUN apk add --no-cache bash
COPY --from=gows /gows /usr/local/bin
WORKDIR /self-hosting-whitespace
COPY self-hosting-whitespace run.bash ./
RUN akaza='ruby -I/akaza/lib /akaza/exe/akaza' && \
    $akaza wsrb src/whitespace.ws.rb > build/whitespace.ws && \
    $akaza wsrb src/fizzbuzz.ws.rb > build/fizzbuzz.ws
# TODO: Cannot find akaza
# RUN rake test
ENTRYPOINT ["./run.bash"]
