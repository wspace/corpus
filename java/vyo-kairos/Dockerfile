FROM eclipse-temurin:8-jdk-jammy AS builder

WORKDIR /kairos
COPY kairos .
RUN ./gradlew shadowJar

FROM eclipse-temurin:8-jre-jammy

WORKDIR /kairos
COPY --from=builder /kairos/build/libs/kairos-1.0.0-all.jar .
ENTRYPOINT ["java", "-jar", "/kairos/kairos-1.0.0-all.jar"]
