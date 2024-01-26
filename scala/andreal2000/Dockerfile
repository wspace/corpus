FROM sbtscala/scala-sbt:eclipse-temurin-jammy-21.0.1_12_1.9.8_3.3.1 AS builder

WORKDIR /WhitespaceEvaluator
COPY WhitespaceEvaluator .
RUN sbt test assembly

FROM eclipse-temurin:21-jre-alpine

WORKDIR /WhitespaceEvaluator
COPY --from=builder ["/WhitespaceEvaluator/target/scala-3.3.1/WhitespaceEvaluator-assembly-0.3.jar", "WhitespaceEvaluator.jar"]
ENTRYPOINT ["java", "-jar", "WhitespaceEvaluator.jar"]
