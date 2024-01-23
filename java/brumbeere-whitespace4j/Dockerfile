FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /Whitespace4j
COPY Whitespace4j .
RUN javac -d bin src/spp/Whitespace4j.java && \
    jar -cfe whitespace4j.jar spp.Whitespace4j -C bin .

FROM eclipse-temurin:21-jre-alpine

WORKDIR /Whitespace4j
COPY --from=builder /Whitespace4j/whitespace4j.jar .
ENTRYPOINT ["java", "-jar", "whitespace4j.jar"]
