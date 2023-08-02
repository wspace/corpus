FROM alpine

RUN apk add --no-cache openjdk8 patch
WORKDIR /space
COPY space jar_task.patch ./
RUN patch -i jar_task.patch && \
    rm jar_task.patch && \
    ./gradlew jar

ENTRYPOINT ["java", "-jar", "build/libs/space.jar"]
