FROM alpine AS builder

RUN apk add make g++
WORKDIR /whitespace-jit
COPY whitespace-jit .
RUN make CXXFLAGS='-O3 -Wall -Wpedantic -static'

FROM scratch

COPY --from=builder /whitespace-jit/compile /
ENTRYPOINT ["/compile"]
