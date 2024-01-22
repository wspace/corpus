FROM sbtscala/scala-sbt:eclipse-temurin-jammy-21.0.1_12_1.9.8_3.3.1 AS builder

WORKDIR /WhitespaceEvaluator
COPY WhitespaceEvaluator .
RUN sbt assembly

FROM eclipse-temurin:21-jre-alpine

WORKDIR /WhitespaceEvaluator
COPY --from=builder ["/WhitespaceEvaluator/target/scala-3.3.1/WhitespaceEvaluator-assembly-0.1.0.jar", "WhitespaceEvaluator.jar"]
ENTRYPOINT ["java", "-jar", "WhitespaceEvaluator.jar"]
