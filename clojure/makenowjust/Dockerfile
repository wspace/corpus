FROM clojure:temurin-18-lein-focal as builder

WORKDIR /whitespace-clj
COPY whitespace-clj .
RUN lein uberjar

FROM eclipse-temurin:18-jre-focal

WORKDIR /whitespace-clj
COPY --from=builder /whitespace-clj/target/whitespace-clj-0.1.0-SNAPSHOT-standalone.jar .
ENTRYPOINT ["java", "-jar", "whitespace-clj-0.1.0-SNAPSHOT-standalone.jar"]
