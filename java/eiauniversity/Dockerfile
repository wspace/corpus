FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /WhiteSpaceInterpreter
COPY WhiteSpaceInterpreter .
RUN javac -d bin src/whitespace/whitespace.java && \
    jar -cfe whitespace.jar whitespace.whitespace -C bin .

FROM eclipse-temurin:21-jre-alpine

WORKDIR /WhiteSpaceInterpreter
COPY --from=builder /WhiteSpaceInterpreter/whitespace.jar /WhiteSpaceInterpreter/src/fichero.bin ./
ENTRYPOINT ["java", "-jar", "whitespace.jar"]
