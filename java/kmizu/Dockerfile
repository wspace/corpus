FROM eclipse-temurin:8-jdk-alpine AS builder

RUN apk add --no-cache apache-ant wget moreutils
RUN mkdir javacc-3.2 && \
    wget https://repo1.maven.org/maven2/net/java/dev/javacc/javacc/3.2/javacc-3.2.jar -O javacc-3.2/javacc.jar

WORKDIR /ws_interpreter
COPY ws_interpreter .
RUN find . -exec sh -c 'iconv -f shift-jis -t utf-8 {} | sponge {}' \; && \
    sed -i -e 's,"C:/development/,"/,' \
           -e 's/value="src"/value="."/' \
           -e 's/source="1.4"/source="8"/' \
           build.xml && \
    ant install

FROM eclipse-temurin:20-jre-alpine

WORKDIR /ws_interpreter
COPY --from=builder /ws/ws.jar .
ENTRYPOINT ["java", "-jar", "ws.jar"]
