FROM perl

RUN cpanm Algorithm::Combinatorics File::Basename Parse::Token::Lite String::Unescape Tie::RegexpHash
WORKDIR /wsm2ws
COPY wsm2ws .
ENTRYPOINT ["perl", "wsm2ws.pl"]
