FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /spacedide
COPY spacedide .
RUN javac -d bin $(find src -name '*.java' ! -path src/ui/SpacedView_Old.java) && \
    cp src/spaced/default_settings bin/spaced/ && \
    jar -cfe spacedide.jar spaced.Spaced -C bin .

FROM eclipse-temurin:21-jre-alpine

WORKDIR /spacedide
COPY --from=builder /spacedide/spacedide.jar .
COPY spacedide/res/examples .
ENTRYPOINT ["java", "-jar", "spacedide.jar"]
