FROM eclipse-temurin:20-jdk-alpine AS builder

RUN apk add --no-cache apache-ant
WORKDIR /whiteSpaceInterpreter
COPY whiteSpaceInterpreter .
RUN ant

FROM eclipse-temurin:20-jre-alpine

WORKDIR /whiteSpaceInterpreter
COPY --from=builder /whiteSpaceInterpreter/dist/WhiteSpaceInterpreter.jar .
COPY whiteSpaceInterpreter/*.ws .
ENTRYPOINT ["java", "-jar", "WhiteSpaceInterpreter.jar"]
