FROM eclipse-temurin:18-jdk-jammy AS builder

WORKDIR /WhiteSpace-Interpreter
COPY WhiteSpace-Interpreter .
RUN javac *.java

FROM eclipse-temurin:18-jre-jammy

WORKDIR /WhiteSpace-Interpreter
COPY --from=builder /WhiteSpace-Interpreter/*.class ./
ENTRYPOINT ["java", "-cp", ".", "Test"]
