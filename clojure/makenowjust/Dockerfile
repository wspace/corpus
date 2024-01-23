FROM clojure:temurin-21-lein-alpine AS builder

WORKDIR /whitespace-clj
COPY whitespace-clj .
RUN lein uberjar

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whitespace-clj
COPY --from=builder /whitespace-clj/target/whitespace-clj-0.1.0-SNAPSHOT-standalone.jar whitespace-clj-0.1.0-SNAPSHOT.jar
ENTRYPOINT ["java", "-jar", "whitespace-clj-0.1.0-SNAPSHOT.jar"]
