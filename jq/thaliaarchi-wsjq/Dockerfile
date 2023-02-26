FROM alpine

RUN apk add --no-cache jq bash
WORKDIR /wsjq
COPY wsjq .
ENTRYPOINT ["./wsjq"]
