FROM eclipse-temurin:20-jdk-alpine AS builder

RUN apk add --no-cache apache-ant wget patch
RUN wget https://dlcdn.apache.org/ant/ivy/2.5.1/apache-ivy-2.5.1-bin.zip && \
    unzip -j apache-ivy-2.5.1-bin.zip apache-ivy-2.5.1/ivy-2.5.1.jar -d /usr/share/java/apache-ant/lib/ && \
    rm -rf apache-ivy-2.5.1-bin.zip

WORKDIR /whitespace4j
COPY whitespace4j jar.patch ./
RUN patch -p1 -i jar.patch && \
    ant test jar

FROM eclipse-temurin:20-jre-alpine

WORKDIR /whitespace4j
COPY --from=builder /whitespace4j/whitespace4j.jar .
ENTRYPOINT ["java", "-jar", "whitespace4j.jar"]
