FROM eclipse-temurin:8-jdk-alpine AS builder

RUN apk add --no-cache bash
WORKDIR /kairos
COPY kairos .
RUN ./gradlew shadowJar

FROM eclipse-temurin:20-jre-alpine

WORKDIR /kairos
COPY --from=builder /kairos/build/libs/kairos-1.0.0-all.jar kairos-1.0.0.jar
ENTRYPOINT ["java", "-jar", "kairos-1.0.0.jar"]
