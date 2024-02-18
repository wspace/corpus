# Fails to run Gradle on alpine due to javax.net.ssl.SSLHandshakeException
FROM eclipse-temurin:8-jdk-jammy AS builder

WORKDIR /space
COPY space jar.patch ./
RUN <<EOF
set -e

apt-get update
apt-get install -y --no-install-recommends patch
rm -rf /var/lib/apt/lists/*

patch -p1 -i jar.patch
./gradlew jar

EOF

FROM eclipse-temurin:21-jre-alpine

WORKDIR /space
COPY --from=builder /space/build/libs/space.jar .
ENTRYPOINT ["java", "-jar", "space.jar"]
