FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /whitespace4j
COPY whitespace4j jar.patch ./
RUN <<EOF
set -e

apk add --no-cache apache-ant wget patch
wget https://dlcdn.apache.org/ant/ivy/2.5.2/apache-ivy-2.5.2-bin.zip
unzip -j apache-ivy-2.5.2-bin.zip apache-ivy-2.5.2/ivy-2.5.2.jar -d /usr/share/java/apache-ant/lib/
rm -rf apache-ivy-2.5.2-bin.zip

patch -p1 -i jar.patch
ant test jar

EOF

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whitespace4j
COPY --from=builder /whitespace4j/whitespace4j.jar .
ENTRYPOINT ["java", "-jar", "whitespace4j.jar"]
