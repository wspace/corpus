FROM eclipse-temurin:20-jdk-alpine AS builder

RUN apk add --no-cache apache-ant
WORKDIR /whitespace
COPY whitespace .
RUN ant all && \
    jar -cfe whitespace.jar org.codeberg.jbanana.whitespace.Whitespace -C classes .

FROM eclipse-temurin:20-jre-alpine

WORKDIR /whitespace
COPY --from=builder /whitespace/whitespace.jar .
COPY /whitespace/examples examples
COPY /whitespace/wstests wstests
COPY run.sh .
ENTRYPOINT ["./run.sh"]
