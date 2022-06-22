FROM clojure:temurin-18-lein-focal as builder

WORKDIR /erasure
COPY erasure .
RUN lein uberjar

FROM eclipse-temurin:18-jre-focal

WORKDIR /erasure
COPY --from=builder /erasure/target/uberjar/erasure-2.0-standalone.jar .
ENTRYPOINT ["java", "-jar", "erasure-2.0-standalone.jar"]
