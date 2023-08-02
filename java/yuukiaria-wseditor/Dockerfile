FROM eclipse-temurin:20-jdk-alpine AS builder

WORKDIR /wseditor
COPY wseditor .
RUN javac -d bin $(find src -name '*.java') && \
    jar -cfe wseditor.jar wseditor.Main -C bin .

FROM eclipse-temurin:20-jre-alpine

WORKDIR /wseditor
COPY --from=builder /wseditor/wseditor.jar .
ENTRYPOINT ["java", "-jar", "wseditor.jar"]
