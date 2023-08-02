FROM clojure:temurin-20-lein-alpine AS builder

WORKDIR /erasure
COPY erasure .
RUN lein uberjar

FROM eclipse-temurin:20-jre-alpine

WORKDIR /erasure
COPY --from=builder /erasure/target/uberjar/erasure-2.0-standalone.jar erasure-2.0.jar
ENTRYPOINT ["java", "-jar", "erasure-2.0.jar"]
