FROM eclipse-temurin:18-jdk-jammy AS builder

WORKDIR /WhiteSpaceInterpreter
COPY WhiteSpaceInterpreter .
RUN javac src/whitespace/whitespace.java

FROM eclipse-temurin:18-jre-jammy

WORKDIR /WhiteSpaceInterpreter
COPY --from=builder /WhiteSpaceInterpreter/src/whitespace/whitespace.class ./whitespace/
COPY --from=builder /WhiteSpaceInterpreter/src/fichero.bin .
ENTRYPOINT ["java", "-cp", ".", "whitespace.whitespace"]
