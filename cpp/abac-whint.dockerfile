FROM alpine as builder

RUN apk add git g++
RUN git clone https://github.com/abac00s/WHINT
WORKDIR /WHINT
RUN g++ -O3 -Wall -o whint whint.cpp

FROM scratch as runner

COPY --from=builder /WHINT/whint /
ENTRYPOINT ["/whint"]
