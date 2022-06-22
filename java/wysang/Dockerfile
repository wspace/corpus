FROM eclipse-temurin:18-jdk-jammy AS builder

WORKDIR /Whitespace-Compiler
COPY Whitespace-Compiler .
RUN curl -O https://www.antlr.org/download/antlr-4.5.2-complete.jar
# whitespaceParser is manually modified, so a freshly-generated ANTLR parser
# will not work:
# RUN java -jar antlr-4.5.2-complete.jar whitespace.g4
RUN javac -cp antlr-4.5.2-complete.jar *.java

FROM eclipse-temurin:18-jre-jammy

WORKDIR /Whitespace-Compiler
COPY --from=builder /Whitespace-Compiler/*.class ./
COPY --from=builder /Whitespace-Compiler/antlr-4.5.2-complete.jar .
COPY --from=builder /Whitespace-Compiler/*.whitespace ./
ENTRYPOINT ["java", "-cp", ".:antlr-4.5.2-complete.jar", "LoadWhitespace"]
