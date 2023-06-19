FROM eclipse-temurin:20-jdk-jammy AS builder

WORKDIR /WhiteSpace-Interpreter
COPY WhiteSpace-Interpreter .
RUN javac *.java

FROM eclipse-temurin:20-jre-jammy

WORKDIR /WhiteSpace-Interpreter
COPY --from=builder /WhiteSpace-Interpreter/*.class ./
ENTRYPOINT ["java", "-cp", ".", "Test"]
