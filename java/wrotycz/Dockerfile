FROM maven:3-eclipse-temurin-20-alpine AS builder

WORKDIR /whitespace-interpreter
COPY whitespace-interpreter .
RUN mvn -Dmaven.compiler.source=1.8 -Dmaven.compiler.target=1.8 test package

FROM eclipse-temurin:20-jre-alpine

WORKDIR /whitespace-interpreter
COPY --from=builder /whitespace-interpreter/target/whitespace-interpreter-1.0-SNAPSHOT.jar .
# Needs Codewars entrypoint
ENTRYPOINT ["java", "-jar", "whitespace-interpreter-1.0-SNAPSHOT.jar"]
