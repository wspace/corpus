FROM maven:3-eclipse-temurin-20 AS builder

WORKDIR /whitespace-interpreter
COPY whitespace-interpreter .
RUN mvn test package
RUN tar xf target/whitespace-interpreter-0.4-bin.tar.gz

FROM eclipse-temurin:20-jre-jammy

WORKDIR /whitespace-interpreter
COPY --from=builder /whitespace-interpreter/whitespace-interpreter-0.4 .
ENTRYPOINT ["java", "-jar", "lib/whitespace.jar"]
