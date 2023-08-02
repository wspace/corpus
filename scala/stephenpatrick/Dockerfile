FROM sbtscala/scala-sbt:eclipse-temurin-jammy-19.0.1_10_1.9.3_3.3.0 AS builder

WORKDIR /whitespace-interpreter
COPY whitespace-interpreter .
RUN mkdir project && \
    echo 'addSbtPlugin("com.eed3si9n" % "sbt-assembly" % "2.1.1")' > project/plugins.sbt && \
    sbt assembly

FROM eclipse-temurin:20-jre-alpine

WORKDIR /whitespace-interpreter
COPY --from=builder ["/whitespace-interpreter/target/scala-2.11/Whitespace Interpreter-assembly-1.0.jar", "whitespace.jar"]
COPY whitespace-interpreter/conf conf
ENTRYPOINT ["java", "-jar", "whitespace.jar"]
