FROM alpine AS builder

RUN apk add git make g++
RUN git clone https://github.com/drebelsky/whitespace-jit
WORKDIR /whitespace-jit
RUN make CXXFLAGS='-O3 -Wall -Wpedantic -static'

FROM scratch

COPY --from=builder /whitespace-jit/compile /
ENTRYPOINT ["/compile"]
