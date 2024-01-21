FROM sbtscala/scala-sbt:eclipse-temurin-jammy-21.0.1_12_1.9.8_3.3.1 AS builder

WORKDIR /whitespace-interpreter
COPY whitespace-interpreter .
RUN mkdir project && \
    echo 'addSbtPlugin("com.eed3si9n" % "sbt-assembly" % "2.1.5")' > project/plugins.sbt && \
    sbt assembly

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whitespace-interpreter
COPY --from=builder ["/whitespace-interpreter/target/scala-2.11/Whitespace Interpreter-assembly-1.0.jar", "whitespace.jar"]
COPY whitespace-interpreter/conf conf
ENTRYPOINT ["java", "-jar", "whitespace.jar"]
