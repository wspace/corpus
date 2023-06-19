FROM clojure:temurin-20-lein-jammy AS builder

WORKDIR /whitespace-clj
COPY whitespace-clj .
RUN lein uberjar

FROM eclipse-temurin:20-jre-jammy

WORKDIR /whitespace-clj
COPY --from=builder /whitespace-clj/target/whitespace-clj-0.1.0-SNAPSHOT-standalone.jar .
ENTRYPOINT ["java", "-jar", "whitespace-clj-0.1.0-SNAPSHOT-standalone.jar"]
