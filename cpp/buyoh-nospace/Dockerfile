FROM alpine AS builder

RUN apk add git make g++ ruby
RUN git clone https://github.com/buyoh/nospace
WORKDIR /nospace
RUN make release CXXFLAGS='-O3 -static'
RUN ./test.rb

FROM scratch

COPY --from=builder /nospace/maicomp /
ENTRYPOINT ["/maicomp"]
