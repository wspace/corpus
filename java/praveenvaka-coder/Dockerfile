FROM maven:3-eclipse-temurin-18

WORKDIR /whitespace-coder
COPY whitespace-coder .
RUN mvn -Dmaven.compiler.source=1.8 -Dmaven.compiler.target=1.8 package
RUN java -cp target/whitespace-coder-1.0-SNAPSHOT.jar Factorial
RUN java -cp target/whitespace-coder-1.0-SNAPSHOT.jar FCTRIZE
RUN java -cp target/whitespace-coder-1.0-SNAPSHOT.jar HelloWorld
