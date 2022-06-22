FROM maven:3-eclipse-temurin-18 AS builder

WORKDIR /whitespace-interpreter
COPY whitespace-interpreter .
RUN mvn test package
RUN tar xf target/whitespace-interpreter-0.4-bin.tar.gz

FROM eclipse-temurin:18-jre-jammy

WORKDIR /whitespace-interpreter
COPY --from=builder /whitespace-interpreter/whitespace-interpreter-0.4 .
ENTRYPOINT ["java", "-jar", "lib/whitespace.jar"]
