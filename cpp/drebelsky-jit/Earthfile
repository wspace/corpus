FROM alpine AS builder

RUN apk add --no-cache make g++
WORKDIR /whitespace-jit
COPY whitespace-jit .
RUN make CXXFLAGS='-O3 -static'

FROM scratch

COPY --from=builder /whitespace-jit/compile /
ENTRYPOINT ["/compile"]
