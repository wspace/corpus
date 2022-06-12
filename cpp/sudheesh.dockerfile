FROM alpine as builder

RUN apk add git g++
RUN git clone https://github.com/sudheesh4/Whitespace-Interpreter-C-
WORKDIR /Whitespace-Interpreter-C-
RUN g++ -O3 -o space space.cpp

FROM scratch as runner

COPY --from=builder /Whitespace-Interpreter-C-/space /
ENTRYPOINT ["/space"]
