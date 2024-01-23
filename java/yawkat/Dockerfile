FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /Whitespace
COPY Whitespace .
RUN javac -d bin $(find src -name '*.java') && \
    jar -cfe whitespace.jar at.yawk.whitspace.Whitespace -C bin .

FROM eclipse-temurin:21-jre-alpine

WORKDIR /Whitespace
COPY --from=builder /Whitespace/whitespace.jar .
ENTRYPOINT ["java", "-jar", "whitespace.jar"]
