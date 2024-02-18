# Fails to run Gradle on alpine due to javax.net.ssl.SSLHandshakeException
FROM eclipse-temurin:8-jdk-jammy AS builder

WORKDIR /kairos
COPY kairos .
RUN ./gradlew shadowJar

FROM eclipse-temurin:21-jre-alpine

WORKDIR /kairos
COPY --from=builder /kairos/build/libs/kairos-1.0.0-all.jar kairos-1.0.0.jar
ENTRYPOINT ["java", "-jar", "kairos-1.0.0.jar"]
