FROM eclipse-temurin:20-jdk-alpine AS builder

WORKDIR /whasm
COPY whasm .
RUN javac -d bin WhAsm.java && \
    jar -cfe whasm.jar WhAsm -C bin .

FROM eclipse-temurin:20-jre-alpine

WORKDIR /whasm
COPY --from=builder /whasm/whasm.jar .
ENTRYPOINT ["java", "-jar", "whasm.jar"]
