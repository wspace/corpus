FROM eclipse-temurin:18-jdk-jammy AS builder

WORKDIR /misc-proj
COPY misc-proj .
RUN javac java/whitespace-interpreter/Whitespace.java

FROM eclipse-temurin:18-jre-jammy

WORKDIR /misc-proj
COPY --from=builder /misc-proj/java/whitespace-interpreter/*.class ./
COPY misc-proj/java/whitespace-polyglot/HelloWorld.java java/whitespace-polyglot/
ENTRYPOINT ["java", "-cp", ".", "Whitespace"]
