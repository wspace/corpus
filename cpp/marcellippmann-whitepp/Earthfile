FROM alpine AS builder

RUN apk add --no-cache make g++
WORKDIR /Whitepp
COPY Whitepp .
RUN make

FROM scratch

COPY --from=builder /Whitepp/bin/White++ /
ENTRYPOINT ["/White++"]
