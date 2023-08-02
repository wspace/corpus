FROM eclipse-temurin:8-jdk-alpine AS builder

RUN apk add --no-cache patch
WORKDIR /space
COPY space jar.patch ./
RUN patch -p1 -i jar.patch && \
    ./gradlew jar

FROM eclipse-temurin:20-jre-alpine

WORKDIR /space
COPY --from=builder /space/build/libs/space.jar .
ENTRYPOINT ["java", "-jar", "space.jar"]
