FROM eclipse-temurin:18-jdk-jammy AS builder

WORKDIR /Whitespace4j
COPY Whitespace4j .
RUN javac src/spp/Whitespace4j.java

FROM eclipse-temurin:18-jre-jammy

WORKDIR /Whitespace4j
COPY --from=builder /Whitespace4j/src/spp/*.class spp/
ENTRYPOINT ["java", "-cp", ".", "spp/Whitespace4j"]
