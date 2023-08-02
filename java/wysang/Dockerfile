FROM eclipse-temurin:20-jdk-alpine AS builder

RUN apk add --no-cache curl
WORKDIR /Whitespace-Compiler
COPY Whitespace-Compiler .
RUN curl -O https://www.antlr.org/download/antlr-4.5.2-complete.jar
# whitespaceParser is manually modified, so a freshly-generated ANTLR parser
# will not work:
# RUN java -jar antlr-4.5.2-complete.jar whitespace.g4
RUN javac -cp antlr-4.5.2-complete.jar *.java
RUN mkdir tmp && \
    (cd tmp && \
      jar -xf ../antlr-4.5.2-complete.jar && \
      rm -rf META-INF && \
      cp ../*.class . && \
      jar -cfe ../whitespaceCompiler.jar LoadWhitespace .) && \
    rm -rf tmp

FROM eclipse-temurin:20-jre-alpine

WORKDIR /Whitespace-Compiler
COPY --from=builder /Whitespace-Compiler/whitespaceCompiler.jar .
COPY --from=builder /Whitespace-Compiler/*.whitespace ./
ENTRYPOINT ["java", "-jar", "whitespaceCompiler.jar"]
