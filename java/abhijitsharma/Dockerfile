FROM maven:3-eclipse-temurin-21-alpine AS builder

WORKDIR /whitespace
COPY whitespace .
RUN sed -i 's/>1.6</>1.8</' pom.xml && \
    mvn package

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whitespace
COPY --from=builder /whitespace/target/whitespace-1.0.jar .
COPY --from=builder /whitespace/docs/*.ws /whitespace/docs/addOneTwo.txt ./
ENTRYPOINT ["java", "-jar", "whitespace-1.0.jar"]
