FROM maven:3-eclipse-temurin-21-alpine AS builder

WORKDIR /whitespace-interpreter
COPY whitespace-interpreter jar.patch ./
RUN apk add --no-cache patch && \
    patch -p1 -i jar.patch && \
    mvn compile test assembly:single

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whitespace-interpreter
COPY --from=builder /whitespace-interpreter/target/whitespace-interpreter-0.4-jar-with-dependencies.jar whitespace-interpreter-0.4.jar
COPY whitespace-interpreter/src/main/pack/example/*.ws whitespace-interpreter/src/main/pack/example/*.szm ./
ENTRYPOINT ["java", "-jar", "whitespace-interpreter-0.4.jar"]
