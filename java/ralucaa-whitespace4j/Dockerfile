FROM alpine

RUN apk add --no-cache openjdk17 apache-ant wget
RUN wget https://dlcdn.apache.org/ant/ivy/2.5.1/apache-ivy-2.5.1-bin.zip && \
    unzip -j apache-ivy-2.5.1-bin.zip apache-ivy-2.5.1/ivy-2.5.1.jar -d /usr/share/java/apache-ant/lib/ && \
    rm -rf apache-ivy-2.5.1-bin.zip

WORKDIR /whitespace4j
COPY whitespace4j .
RUN ant

ENTRYPOINT ["java", "-cp", "bin:lib/junit-4.8.2.jar:lib/log4j-1.2.17.jar", "uk.ac.glasgow.whitespace4j.Main"]
