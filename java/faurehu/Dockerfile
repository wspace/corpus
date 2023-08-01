FROM alpine

RUN apk add --no-cache openjdk17 apache-ant

WORKDIR /whiteSpaceInterpreter
COPY whiteSpaceInterpreter .
RUN ant

ENTRYPOINT ["java", "-jar", "dist/WhiteSpaceInterpreter.jar"]
