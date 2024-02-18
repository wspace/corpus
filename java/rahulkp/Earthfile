FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /WhiteSpace-Interpreter
COPY WhiteSpace-Interpreter .
RUN javac -d bin *.java && \
    jar -cfe WhiteSpace.jar Test -C bin .

FROM eclipse-temurin:21-jre-alpine

WORKDIR /WhiteSpace-Interpreter
COPY --from=builder /WhiteSpace-Interpreter/WhiteSpace.jar .
COPY WhiteSpace-Interpreter/*.ws .
ENTRYPOINT ["java", "-jar", "WhiteSpace.jar"]
