FROM erlang

WORKDIR /mrwhite
COPY mrwhite run.sh ./
COPY mrwhite.erl src/
RUN erl -noshell \
        -s leex file priv/from_text_lexer \
        -s yecc file priv/from_text_parser \
        -s leex file priv/from_whitespace_lexer \
        -s yecc file priv/from_whitespace_parser \
        -s erlang halt && \
    mv priv/*.erl src/
RUN rebar3 compile
RUN rebar3 eunit

ENTRYPOINT ["./run.sh"]
