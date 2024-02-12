FROM maven:3-eclipse-temurin-21-alpine AS builder

WORKDIR /whitespace-coder
COPY whitespace-coder .
RUN mvn -Dmaven.compiler.source=1.8 -Dmaven.compiler.target=1.8 package && \
    java -cp target/whitespace-coder-1.0-SNAPSHOT.jar Factorial && \
    java -cp target/whitespace-coder-1.0-SNAPSHOT.jar FCTRIZE && \
    java -cp target/whitespace-coder-1.0-SNAPSHOT.jar HelloWorld

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whitespace-coder
COPY --from=builder /whitespace-coder/target/whitespace-coder-1.0-SNAPSHOT.jar /whitespace-coder/*.ws .
ENTRYPOINT ["java", "-cp", "whitespace-coder-1.0-SNAPSHOT.jar"]
