FROM eclipse-temurin:18-jdk-jammy AS builder

WORKDIR /whasm
COPY whasm .
RUN javac WhAsm.java

FROM eclipse-temurin:18-jre-jammy

WORKDIR /whasm
COPY --from=builder /whasm/*.class ./
ENTRYPOINT ["java", "-cp", ".", "WhAsm"]
