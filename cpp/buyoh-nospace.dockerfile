FROM alpine

RUN apk add git make g++ ruby
WORKDIR /home
RUN git clone https://github.com/buyoh/nospace
WORKDIR /home/nospace
RUN make release
RUN ./test.rb
RUN test -f /home/nospace/maicomp
