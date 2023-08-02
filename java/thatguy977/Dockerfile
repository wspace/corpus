FROM eclipse-temurin:20-jdk-alpine AS builder

WORKDIR /whitespace-interpreter
COPY misc-proj/java/whitespace-interpreter .
RUN javac -d bin Whitespace.java && \
    jar -cfe whitespace.jar Whitespace -C bin .

FROM eclipse-temurin:20-jre-alpine

WORKDIR /whitespace-interpreter
COPY --from=builder /whitespace-interpreter/whitespace.jar .
COPY misc-proj/java/whitespace-polyglot/HelloWorld.java .
ENTRYPOINT ["java", "-jar", "whitespace.jar"]
