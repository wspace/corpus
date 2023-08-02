FROM sbtscala/scala-sbt:eclipse-temurin-jammy-11.0.17_8_1.9.3_3.3.0 AS builder

WORKDIR /whitespace-parser
COPY whitespace-parser update.patch ./
RUN patch -p1 -i update.patch && \
    sbt test assembly

FROM eclipse-temurin:20-jre-alpine

WORKDIR /whitespace-parser
COPY --from=builder /whitespace-parser/target/scala-2.12/whitespace-parser-assembly-0.1.jar whitespace-parser-0.1.jar
ENTRYPOINT ["java", "-jar", "whitespace-parser-0.1.jar"]
