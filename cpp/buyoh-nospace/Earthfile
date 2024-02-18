FROM alpine AS builder

RUN apk add --no-cache make g++ ruby
WORKDIR /nospace
COPY nospace .
RUN make release CXXFLAGS='-O3 -static'
RUN ./test.rb

FROM scratch

COPY --from=builder /nospace/maicomp /
ENTRYPOINT ["/maicomp"]
