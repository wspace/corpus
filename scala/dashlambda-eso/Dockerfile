# Fails with what looks to be an ICE:
# - eclipse-temurin-jammy-21.0.1_12_1.9.7_3.3.1
# - eclipse-temurin-jammy-21.0.1_12_1.9.8_3.3.1
# Works:
# - eclipse-temurin-jammy-19.0.1_10_1.9.5_3.3.1
FROM sbtscala/scala-sbt:eclipse-temurin-jammy-20.0.2_9_1.9.6_3.3.1 AS builder

WORKDIR /Eso
COPY Eso .
RUN sbt test assembly

FROM eclipse-temurin:21-jre-alpine

WORKDIR /Eso
COPY --from=builder /Eso/target/scala-2.13/Eso-2.6.1.jar .
COPY Eso/testResources samples
ENTRYPOINT ["java", "-jar", "Eso-2.6.1.jar"]
