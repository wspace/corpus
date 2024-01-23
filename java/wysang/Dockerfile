FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /Whitespace-Compiler
COPY Whitespace-Compiler .
RUN <<EOF
set -e

apk add --no-cache curl
curl -O https://www.antlr.org/download/antlr-4.5.2-complete.jar

# whitespaceParser is manually modified, so a freshly-generated ANTLR parser
# will not work:
# java -jar antlr-4.5.2-complete.jar whitespace.g4

javac -cp antlr-4.5.2-complete.jar *.java

# Merge whitespaceCompiler and ANTLR jars:
mkdir tmp
cd tmp
jar -xf ../antlr-4.5.2-complete.jar
rm -rf META-INF
cp ../*.class .
jar -cfe ../whitespaceCompiler.jar LoadWhitespace .
cd ..
rm -rf tmp

EOF

FROM eclipse-temurin:21-jre-alpine

WORKDIR /Whitespace-Compiler
COPY --from=builder /Whitespace-Compiler/whitespaceCompiler.jar .
COPY --from=builder /Whitespace-Compiler/*.whitespace ./
ENTRYPOINT ["java", "-jar", "whitespaceCompiler.jar"]
