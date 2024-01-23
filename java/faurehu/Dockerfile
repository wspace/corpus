FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /whiteSpaceInterpreter
COPY whiteSpaceInterpreter .
RUN apk add --no-cache apache-ant && \
    ant

FROM eclipse-temurin:21-jre-alpine

WORKDIR /whiteSpaceInterpreter
COPY --from=builder /whiteSpaceInterpreter/dist/WhiteSpaceInterpreter.jar .
COPY whiteSpaceInterpreter/*.ws .
ENTRYPOINT ["java", "-jar", "WhiteSpaceInterpreter.jar"]
