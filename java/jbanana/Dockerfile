FROM alpine AS builder

RUN apk add --no-cache openjdk17 apache-ant

WORKDIR /whitespace
COPY whitespace run.sh ./
RUN ant

ENTRYPOINT ["./run.sh"]
