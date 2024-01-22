FROM eclipse-temurin:8-jdk-jammy AS builder

RUN apt-get update && \
    apt-get install -y --no-install-recommends patch && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /space
COPY space jar.patch ./
RUN patch -p1 -i jar.patch && \
    ./gradlew jar

FROM eclipse-temurin:21-jre-alpine

WORKDIR /space
COPY --from=builder /space/build/libs/space.jar .
ENTRYPOINT ["java", "-jar", "space.jar"]
