FROM eclipse-temurin:18-jdk-jammy AS builder

WORKDIR /whiteSpaceInterpreter
COPY whiteSpaceInterpreter/src .
RUN javac whitespaceinterpreter/*.java

FROM eclipse-temurin:18-jre-jammy

WORKDIR /whiteSpaceInterpreter
COPY --from=builder /whiteSpaceInterpreter/whitespaceinterpreter/*.class whitespaceinterpreter/
COPY whiteSpaceInterpreter/*.ws ./
ENTRYPOINT ["java", "-cp", ".", "whitespaceinterpreter.WhiteSpaceInterpreter"]
