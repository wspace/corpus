FROM perl:slim

RUN cpanm Acme::Bleach
WORKDIR /madness
COPY madness .
ENTRYPOINT ["perl", "whitespace_bleached.pl"]
