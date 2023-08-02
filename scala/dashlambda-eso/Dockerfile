FROM sbtscala/scala-sbt:eclipse-temurin-jammy-19.0.1_10_1.9.3_3.3.0 AS builder

WORKDIR /Eso
COPY Eso .
RUN sbt test assembly

FROM eclipse-temurin:20-jre-alpine

WORKDIR /Eso
COPY --from=builder /Eso/target/scala-2.13/Eso-2.6.1.jar .
COPY Eso/testResources samples
ENTRYPOINT ["java", "-jar", "Eso-2.6.1.jar"]
