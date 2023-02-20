FROM alpine AS builder

RUN apk add g++
WORKDIR /wsintercpp
COPY wsintercpp .
RUN g++ -O3 -static -o inter inter.cpp

FROM scratch

COPY --from=builder /wsintercpp/inter /
ENTRYPOINT ["/inter"]
