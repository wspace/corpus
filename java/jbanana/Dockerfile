FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /whitespace
COPY whitespace .
RUN apk add --no-cache apache-ant && \
    ant all && \
    jar -cfe whitespace.jar org.codeberg.jbanana.whitespace.Whitespace -C classes .

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whitespace
COPY --from=builder /whitespace/whitespace.jar .
COPY /whitespace/examples examples
COPY /whitespace/wstests wstests
COPY run.sh .
ENTRYPOINT ["./run.sh"]
